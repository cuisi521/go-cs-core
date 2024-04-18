// package conv
// @Author cuisi
// @Date 2024/4/18 15:48:00
// @Desc
package conv

import (
	"strings"
)

// CaseCamel converts a string to CamelCase
func CaseCamel(s string, next bool) string {
	n := ""
	for i := 0; i < len(s); i++ {
		if next {
			n += strings.ToUpper(string(s[i]))
		} else {
			if string(s[i]) != "_" && string(s[i]) != "-" {
				n += string(s[i])
			}
		}
		if string(s[i]) == "_" || string(s[i]) == "-" {
			next = true
		} else {
			next = false
		}

	}

	return n
}
