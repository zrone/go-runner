package task

import (
	"awesome-runner/src/cache"
	"awesome-runner/src/config"
	"awesome-runner/src/logr"
	"awesome-runner/src/sql"
	"awesome-runner/types"
	"github.com/RichardKnop/machinery/v2/tasks"
)

func TaskErrorHandle(task *tasks.Signature, err error) {
	switch err.Error() {
	case types.NOTIFICATION_WORK_SERVER:
		var taskRecord types.TaskLog
		sql.GetLiteInstance().Take(&taskRecord, "uuid = ?", task.UUID)
		if taskRecord != (types.TaskLog{}) {
			sql.GetLiteInstance().Model(&taskRecord).Update("State", `FAILURE`)
		}

		// 记录失败任务
		logr.Logrus.Errorf("Err: %s, TaskId: %s", err.Error(), task.UUID)

		pipeline := cache.GetReidsInstance().Pipeline()
		cache.GetReidsInstance().Do(logr.Ctx, "select", config.Cnf.QueueDb)
		if signature, err := logr.JSON.Marshal(task.Args); err == nil {
			cache.GetReidsInstance().HSet(logr.Ctx, "task:failed:"+types.NOTIFICATION_WORK_SERVER, task.UUID, signature)
		}
		if _, err := pipeline.Exec(logr.Ctx); err != nil {
			// ..
			logr.Logrus.Errorf("Pipeline Err: %s", err.Error())
		}
	default:
		logr.Logrus.Errorf("Unknown: %s", err.Error())
	}
}
