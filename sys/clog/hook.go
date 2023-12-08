package clog

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

// Hook 写文件的Logrus Hook
type Hook struct {
	W          LoggerInterface
	Level      logrus.Level
	Mod        bool
	LineNumber bool
}

func NewHook(file string, level logrus.Level, params ...bool) (f *Hook) {
	w := NewFileWriter()
	config := fmt.Sprintf(`{"filename":"%s","maxdays":30}`, file)
	err := w.Init(config)
	if err != nil {
		return nil
	}
	var mod, lineNumber bool = false, false
	for i, v := range params {
		if i == 0 {
			mod = v
		} else if i == 1 {
			lineNumber = v
		}
	}

	return &Hook{w, level, mod, lineNumber}
}

// Fire 实现Hook的Fire接口
func (hook *Hook) Fire(entry *logrus.Entry) (err error) {
	if hook.Mod && hook.Level != entry.Level {
		return nil
	}
	message, err := getMessage(entry, hook.LineNumber)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	switch entry.Level {
	case logrus.PanicLevel:
		fallthrough
	case logrus.FatalLevel:
		fallthrough
	case logrus.ErrorLevel:
		return hook.W.WriteMsg(fmt.Sprintf("[ERROR] %s", message), LevelError)
	case logrus.WarnLevel:
		return hook.W.WriteMsg(fmt.Sprintf("[WARN] %s", message), LevelWarn)
	case logrus.InfoLevel:
		return hook.W.WriteMsg(fmt.Sprintf("[INFO] %s", message), LevelInfo)
	case logrus.DebugLevel:
		return hook.W.WriteMsg(fmt.Sprintf("[DEBUG] %s", message), LevelDebug)
	default:
		return nil
	}
}

// Levels 实现Hook的Levels接口
func (hook *Hook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
		logrus.DebugLevel,
	}
}

func getMessage(entry *logrus.Entry, lineNumber bool) (message string, err error) {
	message = message + fmt.Sprintf("%s ", entry.Message)
	// file, lineNumber := GetCallerIgnoringLogMulti(2)
	// if file != "" {
	// 	sep := fmt.Sprintf("%s/src/", os.Getenv("GOPATH"))
	// 	fileName := strings.Split(file, sep)
	// 	if len(fileName) >= 2 {
	// 		file = fileName[1]
	// 	}
	// }
	// message = fmt.Sprintf("%s:%d ", file, lineNumber) + message
	//
	if entry.HasCaller() {
		if !lineNumber {
			message = fmt.Sprintf("%s", entry.Caller.Function) + " " + message
		} else {
			message = fmt.Sprintf("%s:%v", entry.Caller.Function, entry.Caller.Line) + " " + message
		}

	}
	for k, v := range entry.Data {
		message = message + fmt.Sprintf("%v:%v ", k, v)
	}
	return
}
