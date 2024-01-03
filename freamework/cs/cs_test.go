// package cs
// @Author cuisi
// @Date 2024/1/3 11:12:00
// @Desc
package cs

import (
	"testing"

	"github.com/cuisi521/go-cs-core/sys/clog"
)

func TestLog(t *testing.T) {
	clog.SetSaveMod(true)
	// 设置日志输出的文件名以及行号
	clog.SetLineNumber(true)
	clog.SetLogPath("./log")
	clog.Install()
	Log().Infof("sssssss%s", "vvvv")
}
