package logr

// package logr
// @Description 日志

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var (
	Logrus *logrus.Entry
	Clog   *Logger
	Ctx    context.Context
	JSON   jsoniter.API
)

func LogInit() {
	// 创建目录
	workdir, _ := os.Getwd()
	if _, err := os.Stat(workdir + `/runtime`); os.IsNotExist(err) {
		os.MkdirAll(workdir+`/runtime/git`, 0766)
		os.MkdirAll(workdir+`/runtime/task`, 0766)
	}

	Ctx = context.Background()
	JSON = jsoniter.ConfigCompatibleWithStandardLibrary

	path := "runtime/process.log"

	writer, _ := rotatelogs.New(
		"runtime/process_%Y%m%d.log",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(writer)
	logrus.SetLevel(logrus.DebugLevel)
	Logrus = logrus.WithContext(Ctx)

	Clog = NewLogger("", 0)
}
