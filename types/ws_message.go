package types

// 消息体
type InputMessage struct {
	UUID string
}

// 返回消息体
type OutputMessage struct {
	Code    int       `json:"code"`
	Data    LogFormat `json:"data,omitempty"`
	Message string    `json:"msg"`
}

type LogFormat struct {
	Level string `json:"level"`
	Msg   string `json:"msg"`
	Time  string `json:"time"`
}
