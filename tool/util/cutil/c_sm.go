// package cutil
// @Author cuisi
// @Date 2024/1/11 15:51:00
// @Desc
package cutil

import (
	"fmt"

	"github.com/tjfoc/gmsm/sm3"
)

// ESM3 SM3加密
func ESM3(content string) (outStr string) {
	h := sm3.New()
	h.Write([]byte(content))
	sum := h.Sum(nil)
	outStr = byteToString(sum)
	return
}

func byteToString(b []byte) string {
	ret := ""
	for i := 0; i < len(b); i++ {
		ret += fmt.Sprintf("%02x", b[i])
	}
	return ret
}
