package logr

// package logr
// @Description 日志

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"time"
)

var (
	Logrus *logrus.Entry
	Clog   *Logger
	Ctx    context.Context
	JSON   jsoniter.API
)

func LogInit() {
	Ctx = context.Background()
	JSON = jsoniter.ConfigCompatibleWithStandardLibrary

	path := "runtime/process.logr"

	writer, _ := rotatelogs.New(
		"runtime/process_%Y%m%d.logr",
		rotatelogs.WithLinkName(path),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(writer)
	logrus.SetLevel(logrus.DebugLevel)
	Logrus = logrus.WithContext(Ctx)

	Clog = NewLogger("", 0)
}
