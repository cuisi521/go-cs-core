// package cache
// @Author cuisi
// @Date 2023/12/7 10:39:00
// @Desc
package cache

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type Cacher interface {
	SetMaxMemory(size string) bool
	Set(key string, value interface{}) (err error)
	SetEX(key string, value interface{}, expiration time.Duration) (err error)
	Get(key string) (result interface{}, err error)
	GetSet(key string, value interface{}) (result interface{}, err error)
	Del(keys ...string) (err error)
	FlushDB() error
}

type memCache struct {
	maxMemorySize                int64
	maxMemorySizeStr             string
	currMemorySize               int64
	values                       map[string]*memCacheValue
	locker                       sync.RWMutex
	cleanExpiredItemTimeInterval time.Duration
}

type memCacheValue struct {
	val        interface{}
	expireTime time.Time
	expire     time.Duration
	size       int64
}

func NewMemCache() Cacher {
	mc := &memCache{
		values:                       make(map[string]*memCacheValue),
		cleanExpiredItemTimeInterval: time.Second,
	}
	go mc.cleanExpiredItem()
	return mc
}

func (mc *memCache) SetMaxMemory(size string) bool {
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	return true
}

func (mc *memCache) Set(key string, val interface{}) error {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	v := &memCacheValue{
		val:  val,
		size: GetValSize(val),
	}
	if _, ok := mc.get(key); ok {
		mc.del(key)
	}
	mc.add(key, v)

	if mc.currMemorySize > mc.maxMemorySize {
		mc.del(key)
		log.Println(fmt.Sprintf("max memory size %s", mc.maxMemorySizeStr))
	}
	return nil
}

func (mc *memCache) SetEX(key string, val interface{}, expire time.Duration) error {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	v := &memCacheValue{
		val:        val,
		expireTime: time.Now().Add(expire),
		expire:     expire,
		size:       GetValSize(val),
	}
	if _, ok := mc.get(key); ok {
		mc.del(key)
	}
	mc.add(key, v)

	if mc.currMemorySize > mc.maxMemorySize {
		mc.del(key)
		log.Println(fmt.Sprintf("max memory size %s", mc.maxMemorySizeStr))
	}
	return nil
}

func (mc *memCache) GetSet(key string, value interface{}) (result interface{}, err error) {
	if !mc.Exists(key) {
		err = mc.Set(key, value)
		result = value
	} else {
		result, err = mc.Get(key)
	}
	return
}

func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}

func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	if ok && tmp != nil {
		mc.currMemorySize -= tmp.size
		delete(mc.values, key)
	}
}

func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currMemorySize += val.size
}

func (mc *memCache) Get(key string) (result interface{}, err error) {
	mc.locker.RLock()
	defer mc.locker.RUnlock()

	mcv, ok := mc.get(key)
	if ok {
		if mcv.expire != 0 && mcv.expireTime.Before(time.Now()) {
			mc.del(key)
			result = nil
			err = fmt.Errorf("不存在key")
			return
		}
		result = mcv.val
		return
	} else {
		err = fmt.Errorf("不存在key")
	}
	return
}

func (mc *memCache) Del(keys ...string) error {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	for _, key := range keys {
		mc.del(key)
	}
	return nil
}

func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	_, ok := mc.values[key]
	return ok
}

func (mc *memCache) FlushDB() error {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.values = make(map[string]*memCacheValue, 0)
	mc.currMemorySize = 0

	return nil
}

func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	defer mc.locker.RUnlock()
	return int64(len(mc.values))
}

func (mc *memCache) cleanExpiredItem() {
	timeTicker := time.NewTicker(mc.cleanExpiredItemTimeInterval)
	defer timeTicker.Stop()
	for {
		select {
		case <-timeTicker.C:
			for key, item := range mc.values {
				if item.expire != 0 && time.Now().After(item.expireTime) {
					mc.locker.Lock()
					mc.del(key)
					mc.locker.Unlock()
				}
			}
		}
	}
}
