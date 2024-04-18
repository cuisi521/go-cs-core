// package conv
// @Author cuisi
// @Date 2024/4/18 15:49:00
// @Desc
package conv

import (
	"fmt"
	"strings"
	"testing"
)

func TestConv(t *testing.T) {
	tableName := "table_name"
	titleCase := toTitleCase(tableName)
	lowerCamelCase := toLowerCamelCase(tableName)
	fmt.Println("大驼峰形式:", titleCase)
	fmt.Println("小驼峰形式:", lowerCamelCase)
	fmt.Println("sss:", convertToCustomCase(tableName, false))
}
func toTitleCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToTitle(word)
	}
	return strings.Join(words, "")
}

func toLowerCamelCase(s string) string {
	s = toTitleCase(s)
	return strings.ToLower(s[:1]) + s[1:]
}

func convertToCustomCase(s string, next bool) string {
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
