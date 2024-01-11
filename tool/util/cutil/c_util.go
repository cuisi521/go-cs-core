// package cutil
// @Author cuisi
// @Date 2024/1/3 10:24:00
// @Desc
package cutil

import (
	"fmt"
)

func Try(try func()) (err error) {
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok {
				err = v
			} else {
				err = fmt.Errorf(`%+v`, exception)
			}
		}
	}()
	try()
	return
}

func TryCatch(try func(), catch ...func(exception error)) {
	defer func() {
		if exception := recover(); exception != nil && len(catch) > 0 {
			if v, ok := exception.(error); ok {
				catch[0](v)
			} else {
				catch[0](fmt.Errorf(`%+v`, exception))
			}
		}
	}()
	try()
}
