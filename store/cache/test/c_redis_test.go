// package test
// @Author cuisi
// @Date 2023/12/5 10:21:00
// @Desc
package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cuisi521/go-cs-core/store/cache"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestHSet(t *testing.T) {
	connErr := install()
	if connErr != nil {
		return
	}

	us := []*User{}
	for i := 0; i < 10; i++ {
		u := User{
			Name: fmt.Sprintf("name%v", i),
			Age:  0,
		}
		us = append(us, &u)
	}

	mu, er := json.Marshal(&us)
	if er != nil {
		fmt.Println("error:", er.Error())
	}
	fmt.Println(string(mu))
	err := cache.Redis().HSet("ks1", "k1", "1", "k2", "2", "k3", "3")
	if err != nil {
		fmt.Println(err.Error())
	}

	err1 := cache.Redis().Set("vs1", mu)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
}

func install() error {
	redisCnf := &cache.Config{
		Address:       "127.0.0.1:9999",
		Db:            0,
		User:          "",
		Pass:          "sowell@123",
		ReadTimeout:   time.Second * 3,
		WriteTimeout:  time.Second * 3,
		WaitTimeout:   time.Second * 10,
		TLS:           false,
		TLSSkipVerify: false,
		TLSConfig:     nil,
		SlaveOnly:     false,
		PoolSize:      50,
	}

	cnf := &cache.RedisCnfs{Alias: "default", Cnf: redisCnf}
	rcnf := make([]*cache.RedisCnfs, 0)
	rcnf = append(rcnf, cnf)
	err := cache.New(rcnf)
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}
