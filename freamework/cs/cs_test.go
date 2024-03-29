// package cs
// @Author cuisi
// @Date 2024/1/3 11:12:00
// @Desc
package cs

import (
	"fmt"
	"testing"

	"github.com/cuisi521/go-cs-core/sys/clog"
)

func TestLog(t *testing.T) {
	clog.SetSaveMod(false)
	// 设置日志输出的文件名以及行号
	clog.SetLineNumber(true)
	clog.SetLogPath("./log")
	clog.Install()
	Log().Infof("sssssss%s", "vvvv")
}

func TestGetType(t *testing.T) {
	var ts interface{}
	ts = []byte("sssss")
	vs := tStruct{Id: "888888"}
	v1 := GetType(ts)
	v2 := GetType(vs)
	fmt.Println(v1)
	fmt.Println(v2)
}

type tStruct struct {
	Id string
}
