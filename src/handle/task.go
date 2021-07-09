package handle

import (
	"awesome-runner/src/logr"
	"awesome-runner/src/queue"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"github.com/RichardKnop/machinery/v2/tasks"
	"github.com/golang-module/carbon"
	"github.com/kataras/iris/v12"
	taskLogrus "github.com/sirupsen/logrus"
	"time"
)

// 重试
func Retry(ctx iris.Context) {
	symbol := ctx.Params().Get("symbol")
	var task types.TaskLog
	sql.GetLiteInstance().Where("symbol = ?", symbol).Take(&task)
	if task == (types.TaskLog{}) {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "Unknown symbol",
		})
		return
	}

	var args []tasks.Arg
	err := logr.JSON.Unmarshal([]byte(task.Args), &args)
	if err != nil {
		ctx.JSON(types.Response{
			Code:    400,
			Message: "fail",
		})
		return
	}

	uuid := logr.SnowFlakeId()
	eta := time.Now().Add(time.Second * 3)

	for i, arg := range args {
		if arg.Name == "UUID" {
			args[i].Value = uuid
			break
		}
	}

	// 发送部署任务
	signature := &tasks.Signature{
		UUID: uuid,
		Name: "call",
		Args: args,
		ETA:  &eta,
		// 重试次数和斐波那契间隔
		// RetryCount:   3,
		// RetryTimeout: 15,
	}

	tl := types.TaskLog{
		Symbol:    task.Symbol,
		Committer: task.Committer,
		Version:   task.Version,
		Uuid:      uuid,
		CreateAt:  carbon.Now().ToTimestamp(),
		Args:      task.Args,
		Type:      task.Type,
	}

	if _, err := queue.MachineryServer.SendTask(signature); err != nil {
		tl.State = "FAILURE"
		sql.GetLiteInstance().Create(&tl)

		taskLogrus.Errorf("Failed task delivered，%v", err)
		logr.Clog.Errorf("Failed task delivered，%v", err)

		ctx.JSON(types.Response{
			Code:    400,
			Message: "fail",
		})
	} else {
		tl.State = "PENDING"
		sql.GetLiteInstance().Create(&tl)

		ctx.JSON(types.Response{
			200,
			"Success task " + uuid + " delivered",
			nil,
		})
	}
}
