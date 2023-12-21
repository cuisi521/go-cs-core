// package clog
// @Author cuisi
// @Date 2023/12/21 10:37:00
// @Desc
package server

import (
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/cuisi521/go-cs-core/sys/clog"
)

var (
	once sync.Once
	l    *ServerLog
	ler  clog.Loger
)

type ServerLog struct {
	// 日志保存路径
	logPath string

	// 输出行号
	lineNumber bool
}

func New(path string) *ServerLog {
	l := new(ServerLog)
	ler = l
	if path != "" {
		l.logPath = path
	} else {
		l.logPath = "./logs"
	}
	// 判断是否存在日志文件夹
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println("[error]", err.Error())
		}
	}
	logPath := clog.NewHook(path+"/server.log", logrus.InfoLevel, false, l.lineNumber)
	logrus.AddHook(logPath)
	return l
}

// SetLogPath 日志保存路径
// @author By Cuisi 2023/10/31 09:56:00
func (l *ServerLog) SetLogPath(path string) {
	l.logPath = path
}

// SetLineNumber 设置输出行号
// @author By Cuisi 2023/10/31 09:57:00
func (l *ServerLog) SetLineNumber(h bool) {
	l.lineNumber = h
}

// Info
// @author By Cuisi 2023/10/31 11:01:00
func Info(i ...interface{}) {
	ler.Info(i...)
}

func (v ServerLog) Info(i ...interface{}) {
	logrus.Info(i...)
}

func Infof(format string, args ...interface{}) {
	ler.Infof(format, args...)
}

func (v ServerLog) Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Error
// @author By Cuisi 2023/10/31 11:04:00
func Error(i ...interface{}) {
	ler.Error(i...)
}

func (v ServerLog) Error(i ...interface{}) {
	logrus.Error(i...)
}

func Errorf(format string, args ...interface{}) {
	ler.Errorf(format, args...)
}

func (v ServerLog) Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Panic
// @author By Cuisi 2023/10/31 14:20:00
func Panic(i ...interface{}) {
	ler.Panic(i...)
}

func (v ServerLog) Panic(i ...interface{}) {
	logrus.Panic(i...)
}

func Panicf(format string, args ...interface{}) {
	ler.Panicf(format, args...)
}

func (v ServerLog) Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// Fatal
// @author By Cuisi 2023/10/31 14:20:00
func Fatal(i ...interface{}) {
	ler.Fatal(i...)
}
func (v ServerLog) Fatal(i ...interface{}) {
	logrus.Fatal(i...)
}

func Fatalf(format string, args ...interface{}) {
	ler.Fatalf(format, args...)
}
func (v ServerLog) Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// Warning
// @author By Cuisi 2023/10/31 14:20:00
func Warning(i ...interface{}) {
	ler.Warning(i...)
}
func (v ServerLog) Warning(i ...interface{}) {
	logrus.Warning(i...)
}

func Warningf(format string, args ...interface{}) {
	ler.Warningf(format, args...)
}
func (v ServerLog) Warningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

// Debug
// @author By Cuisi 2023/10/31 14:21:00
func Debug(i ...interface{}) {
	ler.Debug(i...)
}
func (v ServerLog) Debug(i ...interface{}) {
	logrus.Debug(i...)
}

func Debugf(format string, args ...interface{}) {
	ler.Debugf(format, args...)
}
func (v ServerLog) Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Trace
// @author By Cuisi 2023/10/31 14:21:00
func Trace(i ...interface{}) {
	ler.Trace(i...)
}
func (v ServerLog) Trace(i ...interface{}) {
	logrus.Trace(i...)
}

func Tracef(format string, args ...interface{}) {
	ler.Tracef(format, args...)
}
func (v ServerLog) Tracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

func Install() {
	var path string = l.logPath
	if path == "" {
		path = "./server"
	}
	// 判断是否存在日志文件夹
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println("[error]", err.Error())
		}
	}
	// logrus.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceQuote:      true,                  // 键值对加引号
		TimestampFormat: "2006-01-02 15:04:05", // 时间格式
		FullTimestamp:   true,
	})

	if l.lineNumber {
		// 设置日志输出的文件名以及行号
		logrus.SetReportCaller(true)
	}

	// 日志集中输出到一个文件
	logPath := clog.NewHook(path+"/server.log", logrus.InfoLevel, false, l.lineNumber)
	logrus.AddHook(logPath)

}
