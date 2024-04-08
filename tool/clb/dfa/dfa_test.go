// package dfa
// @Author cuisi
// @Date 2024/4/2 15:12:00
// @Desc
package dfa

import (
	"fmt"
	"testing"
)

func TestWord(t *testing.T) {
	words := []string{
		"你@#$%好",
		"中国@#$%",
	}
	df := NewDFAMather()
	df.Build(words)
	s, r := df.Match("太阳，你$好s中    国")
	fmt.Println(s)
	fmt.Println(r)
}
