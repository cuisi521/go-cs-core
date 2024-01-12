// package crand
// @Author cuisi
// @Date 2024/1/12 09:53:00
// @Desc
package crand

import (
	"math/rand"
	"time"
)

// RS  生成随机字符串
// charsets 默认所有可能的字符串包括特殊字符
func RS(length int, charsets ...string) string {
	var charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+_-*&^%$#@!" // 所有可能的字符集合
	if len(charsets) > 0 {
		charset = charsets[0]
	}
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// RD 生成随机数[0-length]
func RD(length int) int {
	rand.Seed(time.Now().UnixNano())
	// 生成0到length之间的随机整数
	return rand.Intn(length)
}
