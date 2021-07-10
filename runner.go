package main

import (
	"awesome-runner/src/cache"
	"awesome-runner/src/config"
	"awesome-runner/src/handle"
	"awesome-runner/src/logr"
	"awesome-runner/src/queue"
	"flag"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

var (
	configFile = flag.String("f", "etc/config.yaml", "the config file")
	sessionID  = "IRISSESSIONID"
	sess       = sessions.New(sessions.Config{
		Cookie:  sessionID,
		Expires: time.Hour,
	})
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
	// 405
	app.OnAnyErrorCode(func(ctx iris.Context) {
		errMessage := map[int]string{
			405: "Method not allowed",
			404: "Not found",
			500: "Server error",
		}

		ctx.JSON(iris.Map{
			"code":    ctx.GetStatusCode(),
			"message": errMessage[ctx.GetStatusCode()],
		})
	})

	app.UseGlobal(sess.Handler())
	app.HandleDir("/web", "./web")
	app.Handle("POST", "/", handle.DeployHandle)
	app.Handle("GET", "/ws", handle.WsHandler)

	// api
	taskRouter := app.Party("/task", SecurityMiddleware)
	{
		// 自动化部署
		taskRouter.Get("/proj/list", handle.ProjList)
		taskRouter.Delete("/proj/{symbol}", handle.ProjDelete)
		taskRouter.Post("/proj", handle.ProjCreate)
		taskRouter.Patch("/proj/{symbol}", handle.ProjUpdate)
		taskRouter.Get("/console/list", handle.ConsoleList)

		// 发布上线单
		taskRouter.Post("/proj/publish", handle.TaskPublish)
		taskRouter.Get("/release/list", handle.ReleaseList)

		// 用户

		taskRouter.Post("/retry/{symbol}", handle.Retry)
		//taskRouter.Post("/console/cancle/{uuid}")
	}
	// api
	userRouter := app.Party("/user")
	{
		userRouter.Get("/list", SecurityMiddleware, handle.UserList)
		userRouter.Post("/handle", SecurityMiddleware, handle.UserCreate)
		userRouter.Patch("/handle/{id:int64}", SecurityMiddleware, handle.UserUpdate)
		userRouter.Delete("/handle/{id:int64}", SecurityMiddleware, handle.UserDelete)

		userRouter.Get("/info/currentUser", SecurityMiddleware, handle.CurrentUser)
		userRouter.Get("/logout", SecurityMiddleware, handle.Logout)
		userRouter.Post("/login/account", handle.LoginAccount)
	}

	app.Listen(config.Cnf.Host + ":" + config.Cnf.Port)
}

// passport
func SecurityMiddleware(ctx iris.Context) {
	session := sessions.Get(ctx)

	isLogin, err := session.GetBoolean("Runner:Login")
	if err != nil || !isLogin {
		ctx.JSON(iris.Map{
			"currentAuthority": "guest",
			"status":           "error",
			"type":             "account",
		})
		return
	}

	ctx.Next()
}
