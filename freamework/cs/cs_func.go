// package cs
// @Author cuisi
// @Date 2024/1/3 11:09:00
// @Desc
package cs

import (
	"github.com/cuisi521/go-cs-core/tool/util/cutil"
)

// Try 异常处理
func Try(try func()) error {
	return cutil.Try(try)
}

// TryCatch 异常处理
func TryCatch(try func(), catch ...func(exception error)) {
	cutil.TryCatch(try, catch...)
}
