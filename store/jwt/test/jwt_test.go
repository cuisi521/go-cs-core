// package jwt
// @Author cuisi
// @Date 2023/12/25 15:03:00
// @Desc
package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/cuisi521/go-cs-core/store/jwt"
	"github.com/cuisi521/go-cs-core/tool/util/conv"
)

type User struct {
	Name   string `json:"name"`
	Pwd    string `json:"pwd"`
	Remark string `json:"remark"`
}

func TestCreateToken(t *testing.T) {
	jwt.RegisterCache()
	var key []byte = []byte("VVabc8yhfushc78jdn_98Ytdd76ddcty")

	// user := User{Name: "cuisi", Pwd: "ves123", Remark: "111"}
	sc := &jwt.StandardClaims{
		UseKey:    conv.Str(time.Now().UnixNano() / 1e6),
		ExpiresAt: time.Now().Add(time.Second * 15).Unix(),
	}
	fmt.Println(sc.ExpiresAt)

	token, err := jwt.CreateToken(key, sc)
	if err != nil {
		fmt.Println("err:", err.Message, err.Code)
	} else {
		fmt.Println(token)
		// time.Sleep(time.Second * 20)
		parseCacheToken(token)
	}

}

func parseCacheToken(token string) {
	var key []byte = []byte("VVabc8yhfushc78jdn_98Ytdd76ddcty")

	j, err := jwt.ParseToken(token, key)
	if err != nil {
		fmt.Println("error:", err.Message, err.Code)
	} else {
		fmt.Println(j)
		fmt.Println(j.StandardClaims)
	}
}
