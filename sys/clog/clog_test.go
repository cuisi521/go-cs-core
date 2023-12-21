package clog

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	// 设置日志输出文件模式,true集中输出，false按级别文件输出
	SetSaveMod(false)
	// 设置日志输出的文件名以及行号
	SetLineNumber(true)
	SetLogPath("./log")
	Install()
	// for i := 0; i < 100; i++ {
	Info("This is an informational message")
	// Error("This is an informational message")
	// }
	Infof("%s 111", fmt.Sprintf("\033[%dm %d \033[0m", 200, 200))
	// WidthFields(logrus.Fields{"animal": "walres"}).Info("A group of walrus emerges from the ocean")
}
