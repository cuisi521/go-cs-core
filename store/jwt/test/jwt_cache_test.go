// package test
// @Author cuisi
// @Date 2023/12/27 13:59:00
// @Desc
package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cuisi521/go-cs-core/store/jwt"
	"github.com/cuisi521/go-cs-core/tool/util/conv"
)

var key []byte = []byte("VVabc8yhfushc78jdn_98Ytdd76ddcty")

func TestCacheCreateToken(t *testing.T) {
	// user := User{Name: "cuisi", Pwd: "ves123", Remark: "111"}
	sc := &jwt.StandardClaims{
		UseKey:    conv.Str(time.Now().UnixNano() / 1e6),
		ExpiresAt: time.Now().Add(time.Second * 15).Unix(),
	}

	token, err := jwt.CreateToken(key, sc)
	if err != nil {
		fmt.Println("err:", err.Message, err.Code)
	} else {
		fmt.Println(token)
		// time.Sleep(time.Second * 20)

	}

}

func parseToken(token string) {

}
