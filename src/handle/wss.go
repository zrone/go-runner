package handle

import (
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"bufio"
	"errors"
	"fmt"
	"github.com/golang-module/carbon"
	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
	"io"
	"net/http"
	"os/exec"
	"sync"
	"time"
)

var (
	wsUpgrader = websocket.Upgrader{
		ReadBufferSize:    4096,
		WriteBufferSize:   4096,
		EnableCompression: true,
		HandshakeTimeout:  5 * time.Second,
		// 处理跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// 接受消息
type Request struct {
	MessageType int
	Type        int
	Payload     *types.InputMessage
}

// 返回消息
type Response struct {
	MessageType int
	Payload     types.OutputMessage
}

// 客户端连接
type wsConnection struct {
	wsSocket *websocket.Conn // 底层websocket
	inChan   chan *Request   // 读队列
	outChan  chan *Response  // 写队列

	mutex     sync.Mutex // Mutex互斥锁，避免重复关闭管道
	isClosed  bool
	closeChan chan byte // 关闭通知
}

func WsHandler(ctx iris.Context) {
	wsSocket, err := wsUpgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		return
	}

	wsConn := &wsConnection{
		wsSocket:  wsSocket,
		inChan:    make(chan *Request, 1000),
		outChan:   make(chan *Response, 1000),
		closeChan: make(chan byte),
		isClosed:  false,
	}

	// 处理器
	go wsConn.procLoop()
	// 读协程
	go wsConn.wsReadLoop()
	// 写协程
	go wsConn.wsWriteLoop()
}

func (wsConn *wsConnection) wsReadLoop() {
	for {
		// 读一个message
		messageType, data, err := wsConn.wsSocket.ReadMessage()
		if err != nil {
			goto error
		}
		var request *Request
		if string(data) == string("ping\n") {
			request = &Request{
				MessageType: messageType,
				Type:        0,
			}
		} else {
			// 解析接受消息
			requestMessage := types.InputMessage{}
			err = logr.JSON.Unmarshal(data, &requestMessage)
			if err != nil {
				logr.Clog.Errorf("#%v#, 消息体不合法", err)
				request = &Request{
					MessageType: messageType,
					Type:        -1, // 消息体不合法
				}
			} else {
				request = &Request{
					MessageType: messageType,
					Type:        1,
					Payload:     &requestMessage,
				}
			}
		}

		// 放入请求队列
		select {
		case wsConn.inChan <- request:
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
	logr.Clog.Debug("closed.")
}

func (wsConn *wsConnection) wsWriteLoop() {
	for {
		select {
		// 取一个应答
		case msg := <-wsConn.outChan:
			resp, _ := logr.JSON.Marshal(msg.Payload)
			// 写给websocket
			if err := wsConn.wsSocket.WriteMessage(msg.MessageType, resp); err != nil {
				goto error
			}
		case <-wsConn.closeChan:
			goto closed
		}
	}
error:
	wsConn.wsClose()
closed:
	logr.Clog.Debug("websocket is closed.")
}

func (wsConn *wsConnection) procLoop() {
	for {
		msg, err := wsConn.wsRead()
		if err != nil {
			break
		}

		err = wsConn.wsWrite(msg)
		if err != nil {
			break
		}
	}
}

func (wsConn *wsConnection) wsWrite(message *Request) error {
	//message.Payload.UUID
	resp := &Response{
		MessageType: message.MessageType,
		Payload:     types.OutputMessage{},
	}

	switch message.Type {
	case -1:
		resp.Payload.Code = 0
		resp.Payload.Data = types.LogFormat{
			Level: "error",
			Msg:   "数据有问题，不是合法的数据",
			Time:  time.Now().String(),
		}
		resp.Payload.Message = "fail"
	case 0:
		resp.Payload.Code = 1
		resp.Payload.Message = "ping success"
	case 1:
		resp.Payload.Code = 1
		resp.Payload.Message = "success"

		var (
			taskRecord types.TaskLog
			wg         sync.WaitGroup
			cmd        *exec.Cmd
		)

		sql.GetLiteInstance().Take(&taskRecord, "uuid = ?", message.Payload.UUID)
		if taskRecord == (types.TaskLog{}) {
			return errors.New("Invalid uuid, websocket closed")
		}

		if taskRecord.State == "PENDING" {
			resp.Payload.Data = types.LogFormat{
				Level: "default",
				Msg:   "PENDING",
				Time:  carbon.Now().Format("Y-m-d H:i:s"),
			}

			wsConn.outChan <- resp
		} else {
			if taskRecord.State == "RUNNING" {
				cmd = exec.Command("/bin/bash", "-c", fmt.Sprintf("tail -n +0 -f runtime/task/"+message.Payload.UUID+".log"))
			} else { // SUCCESS FAILURE
				cmd = exec.Command("/bin/bash", "-c", fmt.Sprintf("cat runtime/task/"+message.Payload.UUID+".log"))
			}

			stdout, _ := cmd.StdoutPipe()
			cmd.Start()

			wg.Add(1)
			go func() {
				defer func() {
					wg.Done()
				}()

				reader := bufio.NewReader(stdout)
				for {
					select {
					// 读等待
					case <-wsConn.closeChan:
						return
					default:
					}

					readString, err := reader.ReadString('\n')
					if err != nil || err == io.EOF {
						break
					}

					var msg types.LogFormat
					err = logr.JSON.Unmarshal([]byte(readString), &msg)
					if err != nil {
						return
					}
					resp.Payload.Data = msg
					wsConn.outChan <- resp
					time.Sleep(time.Millisecond * 1)
				}
			}()

			wg.Wait()
			cmd.Wait()
		}
	}

	select {
	case wsConn.outChan <- resp:
	case <-wsConn.closeChan:
		return errors.New("websocket closed")
	}
	return nil
}

func (wsConn *wsConnection) wsRead() (*Request, error) {
	select {
	// 读等待
	case msg := <-wsConn.inChan:
		return msg, nil
	case <-wsConn.closeChan:
	}
	return nil, errors.New("websocket closed")
}

func (wsConn *wsConnection) wsClose() {
	err := wsConn.wsSocket.Close()
	if err != nil {
		return
	}

	wsConn.mutex.Lock()
	defer wsConn.mutex.Unlock()
	if !wsConn.isClosed {
		wsConn.isClosed = true
		close(wsConn.closeChan)
	}
}
