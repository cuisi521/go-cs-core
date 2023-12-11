package ccfg

import (
	"fmt"
	"testing"
)

func TestCfg(t *testing.T) {
	// 路径
	SetCfgPath("./testData")
	// 名称
	SetCfgName("test")
	// 类型
	SetCfgType("yaml")
	// 读取
	Install()
	fmt.Println(Get("server.address")) // map[port:3306 url:127.0.0.1] .0.0.1

	fmt.Println(SysCnf().Database["default"]["default"].Link)

}
