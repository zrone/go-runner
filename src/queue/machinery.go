package queue

import (
	config2 "awesome-runner/src/config"
	"github.com/RichardKnop/machinery/v2"
	redisbackend "github.com/RichardKnop/machinery/v2/backends/redis"
	redisbroker "github.com/RichardKnop/machinery/v2/brokers/redis"
	"github.com/RichardKnop/machinery/v2/config"
	eagerlock "github.com/RichardKnop/machinery/v2/locks/eager"
	"sync"
)

var (
	MachineryServer *machinery.Server
	MachineryWork   *machinery.Worker
)

func initMachinery() {
	cnf := &config.Config{
		DefaultQueue: "task",
		//ResultBackend:   FetchConfig("redisSchema") + FetchConfig("redisDNS") + "/" + strconv.Itoa(db),
		//ResultsExpireIn: 86400 * 30, // 30 天
		Redis: &config.RedisConfig{
			MaxIdle:                3,
			MaxActive:              3,
			IdleTimeout:            240,
			Wait:                   true,
			ReadTimeout:            15,
			WriteTimeout:           15,
			ConnectTimeout:         15,
			NormalTasksPollPeriod:  1000,
			DelayedTasksPollPeriod: 500,
			DelayedTasksKey:        "task",
		},
	}

	broker := redisbroker.NewGR(cnf, []string{config2.Cnf.QueueDNS}, config2.Cnf.QueueDb)
	backend := redisbackend.NewGR(cnf, []string{config2.Cnf.QueueDNS}, config2.Cnf.QueueDb)
	lock := eagerlock.New()
	MachineryServer = machinery.NewServer(cnf, broker, backend, lock)

	// Register tasks
	tasksMap := map[string]interface{}{
		// "call": CallForCall,
	}

	if err := MachineryServer.RegisterTasks(tasksMap); err != nil {
		return
	}

	MachineryWork = MachineryServer.NewWorker("task", config2.Cnf.WorkNumber)
	// MachineryWork.SetErrorHandler(ErrorWorkHandle)
}

func StoreMachinery() {
	if MachineryServer == nil {
		var once sync.Once
		once.Do(func() {
			initMachinery()
		})
	}
}
