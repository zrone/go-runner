// Copyright (c) 2017-2020 The Elastos Foundation
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.
//

package logr

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"
)

// Level is the level at which a logger is configured.  All messages sent
// to a level which is below the current level are filtered.
type Level uint32

// Level constants.
const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelOff
)

const (
	Red    = "1;31"
	Green  = "1;32"
	Yellow = "1;33"
	Pink   = "1;35"
	Cyan   = "1;36"
)

func Color(code, msg string) string {
	return fmt.Sprintf("\033[%sm%s\033[m", code, msg)
}

const (
	debugLog uint8 = iota
	infoLog
	warnLog
	errorLog
	fatalLog
	disableLog
)

const (
	KBSize = 1024 << (10 * iota)
	MBSize
	GBSize
)

var (
	levels = []string{
		debugLog:   Color(Green, "[DBG]"),
		infoLog:    Color(Pink, "[INF]"),
		warnLog:    Color(Yellow, "[WRN]"),
		errorLog:   Color(Red, "[ERR]"),
		fatalLog:   Color(Red, "[FAT]"),
		disableLog: "DISABLED",
	}
	levelType = []string{
		debugLog:   Green,
		infoLog:    Pink,
		warnLog:    Yellow,
		errorLog:   Red,
		fatalLog:   Red,
		disableLog: "DISABLED",
	}
	Stdout = os.Stdout
)

const (
	calldepth = 2

	defaultPerLogFileSize int64 = 20 * MBSize
	defaultLogsFolderSize int64 = 5 * GBSize
)

var logger *Logger

func levelName(level uint8) string {
	if int(level) >= len(levels) {
		return fmt.Sprintf("LEVEL%d", level)
	}
	return levels[int(level)]
}

type Logger struct {
	level  uint8 // The log print level
	writer io.Writer
	logger *log.Logger
}

func NewLogger(outputPath string, level uint8) *Logger {
	return &Logger{
		level:  level,
		writer: os.Stdout,
		logger: log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds),
	}
}

func NewDefault(path string, level uint8) *Logger {
	logger = NewLogger(path, level)
	return logger
}

func (l *Logger) Writer() io.Writer {
	return l.writer
}

func (l *Logger) Output(level uint8, a ...interface{}) {
	if l.level <= level {
		//a = append([]interface{}{levelName(level)}, a...)
		//l.logger.Output(calldepth, fmt.Sprintln(a...))
		fmt.Println(Color(levelType[int(level)], fmt.Sprintf("%s", a...)))
	}
}

func (l *Logger) Outputf(level uint8, format string, v ...interface{}) {
	if l.level <= level {
		//v = append([]interface{}{levelName(level), "GID", goid()}, v...)
		//l.logger.Output(calldepth, fmt.Sprintf("%s %s %s, "+format+"\n", v...))
		fmt.Println(Color(levelType[int(level)], fmt.Sprintf(format, v...)))
	}
}

func (l *Logger) Debug(a ...interface{}) {
	l.Output(debugLog, a...)
}

func (l *Logger) Debugf(format string, a ...interface{}) {
	l.Outputf(infoLog, format, a...)
}

func (l *Logger) Info(a ...interface{}) {
	l.Output(infoLog, a...)
}

func (l *Logger) Infof(format string, a ...interface{}) {
	l.Outputf(infoLog, format, a...)
}

func (l *Logger) Warn(a ...interface{}) {
	l.Output(warnLog, a...)
}

func (l *Logger) Warnf(format string, a ...interface{}) {
	l.Outputf(warnLog, format, a...)
}

func (l *Logger) Error(a ...interface{}) {
	if l.level <= errorLog {
		l.Output(errorLog, a...)
	}
}

func (l *Logger) Errorf(format string, a ...interface{}) {
	l.Outputf(errorLog, format, a...)
}

func (l *Logger) Fatal(a ...interface{}) {
	l.Output(fatalLog, a...)
}

func (l *Logger) Fatalf(format string, a ...interface{}) {
	l.Outputf(fatalLog, format, a...)
}

// Level returns the current logging level.
func (l *Logger) Level() Level {
	return Level(l.level)
}

// SetLevel changes the logging level to the passed level.
func (l *Logger) SetLevel(level Level) {
	l.level = uint8(level)
}

func Debug(a ...interface{}) {
	logger.Debug(a...)
}

func Debugf(format string, a ...interface{}) {
	logger.Debugf(format, a...)
}

func Info(a ...interface{}) {
	logger.Info(a...)
}

func Warn(a ...interface{}) {
	logger.Warn(a...)
}

func Error(a ...interface{}) {
	logger.Error(a...)
}

func Fatal(a ...interface{}) {
	logger.Fatal(a...)
}

func Infof(format string, a ...interface{}) {
	logger.Infof(format, a...)
}

func Warnf(format string, a ...interface{}) {
	logger.Warnf(format, a...)
}

func Errorf(format string, a ...interface{}) {
	logger.Errorf(format, a...)
}

func Fatalf(format string, a ...interface{}) {
	logger.Fatalf(format, a...)
}

func SetPrintLevel(level uint8) {
	logger.SetLevel(Level(level))
}

// goid returns the current goroutine id.
func goid() string {
	var buf [32]byte
	n := runtime.Stack(buf[:], false)
	fields := strings.Fields(string(buf[:n]))
	if len(fields) <= 1 {
		return ""
	}
	return fields[1]
}
