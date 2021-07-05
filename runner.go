package main

import (
	"awesome-runner/src/cache"
	"awesome-runner/src/config"
	"awesome-runner/src/handle"
	"awesome-runner/src/logr"
	"awesome-runner/src/queue"
	"flag"
	"github.com/kataras/iris/v12"
)

var (
	configFile = flag.String("f", "etc/config.yaml", "the config file")
)

func main() {
	// log
	logr.LogInit()
	logr.Clog.Debug(`
   _____         _____                             
  / ____|       |  __ \                            
 | |  __  ___   | |__) |   _ _ __  _ __   ___ _ __ 
 | | |_ |/ _ \  |  _  / | | | '_ \| '_ \ / _ \ '__|
 | |__| | (_) | | | \ \ |_| | | | | | | |  __/ |   
  \_____|\___/  |_|  \_\__,_|_| |_|_| |_|\___|_|
`)

	flag.Parse()
	// 加载配置
	config.StoreConfig(configFile)
	// 队列初始化
	queue.StoreMachinery()

	go func() {
		if err := queue.MachineryWork.Launch(); err != nil {
			logr.Logrus.Errorf("Work 进程异常，%v", err)
			return
		}
	}()

	defer func() {
		if err := cache.GetReidsInstance().Close(); err != nil {
			logr.Logrus.Error(err)
		}
	}()

	app := iris.New()
	app.Handle("POST", "/", handle.DeployHandle)
	app.Handle("GET", "/ws", handle.WsHandler)
	app.Listen(":" + config.Cnf.Port)

	//sql.GetLiteInstance().Create(&types.InternalDeploy{
	//	Symbol: "demo",
	//	Secret: "9b84cce730b4ddec467ac439f0ec5c3dec08a743",
	//	Path:   "/www/wwwroot/demo",
	//	Auth: types.Authentication{
	//		1,
	//		"root",
	//		"localhost",
	//		22,
	//		"bluestone",
	//	},
	//})
}
