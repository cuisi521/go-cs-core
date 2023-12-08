package clog

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	// 设置日志输出文件模式,true集中输出，false按级别文件输出
	SetSaveMod(false)
	// 设置日志输出的文件名以及行号
	SetLineNumber(false)
	SetLogPath("./log")
	Install()
	Info("This is an informational message")
	WidthFields(logrus.Fields{"animal": "walres"}).Info("A group of walrus emerges from the ocean")
}
