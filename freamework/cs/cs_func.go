// package cs
// @Author cuisi
// @Date 2024/1/3 11:09:00
// @Desc
package cs

import (
	"reflect"

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

// GetType 获取类型
func GetType(value interface{}) string {
	t := reflect.TypeOf(value)
	// 根据类型判断
	if t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8 {
		return "byte"
	} else if t.Kind() == reflect.String {
		return "string"
	} else if t.Kind() == reflect.Interface {
		return "interface"
	} else if t.Kind() == reflect.Struct {
		return "struct"
	} else {
		return t.Kind().String()
	}
}
