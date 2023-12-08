// package test
// @Author cuisi
// @Date 2023/12/7 11:21:00
// @Desc
package test

import (
	"fmt"
	"testing"
	"time"

	"go-cs-orm/store/cache"
)

func TestCatche(t *testing.T) {
	testData := []struct {
		key    string
		val    interface{}
		expire time.Duration
	}{
		{"baer", 678, time.Second * 10},
		{"hrws", false, time.Second * 11},
		{"gddfas", true, time.Second * 12},
		{"rwe", map[string]interface{}{"a": 3, "b": false}, time.Second * 13},
		{"rqew", "fsdfas", time.Second * 14},
		{"fsdew", "这里是字符串这里是字符串这里是字符串", time.Second * 15},
	}

	c := cache.NewMemCache()
	c.SetMaxMemory("10MB")
	for _, item := range testData {
		c.SetEX(item.key, item.val, item.expire)
		val, err := c.Get(item.key)
		if err != nil {
			t.Error("缓存取值失败")
		}
		if item.key != "rwe" && val != item.val {
			t.Error("缓存取值数据与预期不一致")
		}
		_, ok1 := val.(map[string]interface{})
		if item.key == "rwe" && !ok1 {
			t.Error("缓存取值数据与预期不一致")
		}
	}
	for _, item := range testData {
		r, e := c.Get(item.key)
		if e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println(r)
		}
	}

	// if int64(len(testData)) != c.Keys() {
	// 	t.Error("缓存数量不一致")
	// }

	c.Del(testData[0].key)
	c.Del(testData[1].key)

	// if int64(len(testData)) != c.Keys()+2 {
	// 	t.Error("缓存数量不一致")
	// }

	time.Sleep(time.Second * 13)
	fmt.Println("===========================")
	for _, item := range testData {
		r, e := c.Get(item.key)
		if e != nil {
			fmt.Println(e.Error())
		} else {
			fmt.Println(r)
		}
	}
	// if c.Keys() != 0 {
	// 	t.Error("缓存清空失败")
	// }
}
