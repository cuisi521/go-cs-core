// package cs
// @Author cuisi
// @Date 2024/1/3 10:45:00
// @Desc
package cs

import (
	"github.com/xormplus/xorm"

	"github.com/cuisi521/go-cs-core/store/cache"
	"github.com/cuisi521/go-cs-core/store/cdb"
	"github.com/cuisi521/go-cs-core/sys/ccfg"
	"github.com/cuisi521/go-cs-core/sys/clog"
)

// DB 返回数据库xorm的对象
func DB(name ...string) *xorm.Engine {
	return cdb.DB(name...)
}

// Log 日志对象
func Log() *clog.Log {
	return clog.GetLog()
}

// Cfg 系统配置文件
func Cfg() *ccfg.Config {
	return ccfg.SysCnf()
}

// Redis engine
func Redis(name ...string) *cache.RedisEngine {
	return cache.Redis(name...)
}

// Memory 对象
func Memory() cache.Cacher {
	return cache.GetMemCache()
}

// AutoCache 根据配置文件自动选择
// 在redis和内存自动选择引擎
func AutoCache(name ...string) (c cache.Cacher) {
	return cache.AutoCache(name...)
}
