// Package clog
// @Author cuisi
// @Date 2023/10/29 16:23:00
// @Desc config
package clog

import (
	"log"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	once sync.Once
	l    *Log
	ler  Loger
)

type Loger interface {
	Info(i ...interface{})
	Error(i ...interface{})
	Panic(i ...interface{})
	Fatal(i ...interface{})
	Warning(i ...interface{})
	Debug(i ...interface{})
	Trace(i ...interface{})
	WidthField(k string, v interface{}) *Entity
	WidthFields(m map[string]interface{}) *Entity
}

type Log struct {
	// 日志保存路径
	logPath string

	// 日志文件保存模式
	// saveMod=true，日志输出到对应的类型文件，否则保存在一个文件
	saveMod bool

	// 输出行号
	lineNumber bool
}

func init() {
	once.Do(func() {
		l = New()
		ler = l
	})
}

func New() *Log {
	l := new(Log)
	return l
}

// SetLogPath 日志保存路径
// @author By Cuisi 2023/10/31 09:56:00
func SetLogPath(path string) {
	l.SetLogPath(path)
}
func (l *Log) SetLogPath(path string) {
	l.logPath = path
}

// SetSaveMod 日志保存模式
// mod=true 日志输出到对应的文件，否则集中输出到一个文件
// @author By Cuisi 2023/10/31 09:55:00
func SetSaveMod(mod bool) {
	l.SetSaveMod(mod)
}
func (l *Log) SetSaveMod(mod bool) {
	l.saveMod = mod
}

// SetLineNumber 设置输出行号
// @author By Cuisi 2023/10/31 09:57:00
func SetLineNumber(h bool) {
	l.SetLineNumber(h)
}
func (l *Log) SetLineNumber(h bool) {
	l.lineNumber = h
}

// Info
// @author By Cuisi 2023/10/31 11:01:00
func Info(i ...interface{}) {
	ler.Info(i...)
}

func (v Log) Info(i ...interface{}) {
	logrus.Info(i...)
}

// Error
// @author By Cuisi 2023/10/31 11:04:00
func Error(i ...interface{}) {
	ler.Error(i...)
}

func (v Log) Error(i ...interface{}) {
	logrus.Error(i...)
}

// Panic
// @author By Cuisi 2023/10/31 14:20:00
func Panic(i ...interface{}) {
	ler.Panic(i...)
}

func (v Log) Panic(i ...interface{}) {
	logrus.Panic(i...)
}

// Fatal
// @author By Cuisi 2023/10/31 14:20:00
func Fatal(i ...interface{}) {
	ler.Fatal(i...)
}
func (v Log) Fatal(i ...interface{}) {
	logrus.Fatal(i...)
}

// Warning
// @author By Cuisi 2023/10/31 14:20:00
func Warning(i ...interface{}) {
	ler.Warning(i...)
}
func (v Log) Warning(i ...interface{}) {
	logrus.Fatal(i...)
}

// Debug
// @author By Cuisi 2023/10/31 14:21:00
func Debug(i ...interface{}) {
	ler.Debug(i...)
}
func (v Log) Debug(i ...interface{}) {
	logrus.Fatal(i...)
}

// Trace
// @author By Cuisi 2023/10/31 14:21:00
func Trace(i ...interface{}) {
	ler.Trace(i...)
}
func (v Log) Trace(i ...interface{}) {
	logrus.Fatal(i...)
}

func (v Log) WidthFields(m map[string]interface{}) *Entity {
	return &Entity{data: m}
}

func (v Log) WidthField(k string, s interface{}) *Entity {
	d := make(map[string]interface{}, 1)
	d[k] = s
	return &Entity{data: d}
}

func Install() {
	var path string = l.logPath
	if path == "" {
		path = "./logs"
	}
	// 判断是否存在日志文件夹
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			log.Println("[error]", err.Error())
		}
	}
	if l.lineNumber {
		// 设置日志输出的文件名以及行号
		logrus.SetReportCaller(true)
	}
	if !l.saveMod {
		// 日志集中输出到一个文件
		logPath := NewHook(path+"/log.log", logrus.InfoLevel, l.saveMod, l.lineNumber)
		logrus.AddHook(logPath)
	} else {
		infoPath := NewHook(path+"/info.log", logrus.InfoLevel, l.saveMod, l.lineNumber)
		errPath := NewHook(path+"/error.log", logrus.ErrorLevel, l.saveMod, l.lineNumber)
		panicPath := NewHook(path+"/panic.log", logrus.PanicLevel, l.saveMod, l.lineNumber)
		fatalPath := NewHook(path+"/fatal.log", logrus.FatalLevel, l.saveMod, l.lineNumber)
		warnPath := NewHook(path+"/warning.log", logrus.WarnLevel, l.saveMod, l.lineNumber)
		debugPath := NewHook(path+"/debug.log", logrus.DebugLevel, l.saveMod, l.lineNumber)
		tracePath := NewHook(path+"/trace.log", logrus.TraceLevel, l.saveMod, l.lineNumber)
		logrus.AddHook(infoPath)
		logrus.AddHook(errPath)
		logrus.AddHook(panicPath)
		logrus.AddHook(fatalPath)
		logrus.AddHook(warnPath)
		logrus.AddHook(debugPath)
		logrus.AddHook(tracePath)
	}

}
