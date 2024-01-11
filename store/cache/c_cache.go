// package cache
// @Author cuisi
// @Date 2024/1/3 14:09:00
// @Desc
package cache

import (
	"github.com/cuisi521/go-cs-core/freamework/cs"
)

func AutoCache(name ...string) (rc Cacher) {

	switch cs.Cfg().Cache.Mod {
	case 1:
		rc = Redis(name...)
	case 2:
		rc = NewMemCache()
		rc.SetMaxMemory("500MB")
	default:
		rc = NewMemCache()
		rc.SetMaxMemory("500MB")
	}
	return
}
