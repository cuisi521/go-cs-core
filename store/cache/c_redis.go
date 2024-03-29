// package credis
// @Author cuisi
// @Date 2023/12/5 09:31:00
// @Desc
package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

var (
	re = make(map[string]*RedisEngine)
)

type RedisEngine struct {
	ctx context.Context
	cnf *Config
	db  *redis.Client
}

// Redis 获取 RedisEngine
func Redis(db ...string) *RedisEngine {
	key := "default"
	if len(db) > 0 {
		key = db[0]
	}
	return re[key]
}

// New 创建redis连接对象
func New(cnf []*RedisCnfs) error {
	var err error
	for _, c := range cnf {
		r := &RedisEngine{cnf: c.Cnf}
		re[c.Alias] = r
		err = r.Connect()
	}
	return err
}

// Connect 连接redis
func (r *RedisEngine) Connect() (err error) {
	r.ctx = context.Background()
	option := &redis.Options{
		Addr:            r.cnf.Address,
		Password:        r.cnf.Pass,
		DB:              r.cnf.Db,
		ConnMaxIdleTime: r.cnf.MaxConnLifetime,
		PoolSize:        r.cnf.PoolSize,
		ReadTimeout:     r.cnf.ReadTimeout,
		WriteTimeout:    r.cnf.WriteTimeout,
	}
	// tls
	if r.cnf.TLS {
		option.TLSConfig = r.cnf.TLSConfig
	}
	timeoutCtx, cancel := context.WithTimeout(context.Background(), r.cnf.WaitTimeout)
	defer cancel()

	r.db = redis.NewClient(option)
	// 验证是否连接到redis服务端
	if err = r.db.Ping(timeoutCtx).Err(); err != nil {
		err = fmt.Errorf("redis 连接失败:%v", err)
		return err
	}
	return
}

// Set 设置key的值
func (r *RedisEngine) Set(key string, value interface{}) (err error) {
	var cv []byte
	cv, err = json.Marshal(value)
	if err != nil {
		return
	}
	err = r.db.Set(r.ctx, key, cv, 0).Err()
	return
}

// SetEX 设置key的值并指定过期时间
func (r *RedisEngine) SetEX(key string, value interface{}, expiration time.Duration) (err error) {
	var cv []byte
	cv, err = json.Marshal(value)
	if err != nil {
		return
	}
	err = r.db.Set(r.ctx, key, cv, expiration).Err()
	return
}

// Get 通过key获取值
func (r *RedisEngine) Get(key string) (result interface{}, err error) {
	result, err = r.db.Get(r.ctx, key).Result()
	return
}

// GetSet 设置新值获取旧值
func (r *RedisEngine) GetSet(key string, value interface{}) (result interface{}, err error) {
	result, err = r.db.GetSet(r.ctx, key, value).Result()
	return
}

// GetOrSetFuncLock 判断key是否存在，不存在调用callBack方法，将返回值插入内存数据库，注意加互斥锁解决并发安全问题
func (r *RedisEngine) GetOrSetFuncLock(key string, callBack Func, expiration time.Duration) (result interface{}, err error) {
	result, err = r.db.Get(r.ctx, key).Result()
	if result == "" || err != nil {
		result, err = callBack()
		if err != nil {
			return
		}
		// t := reflect.TypeOf(result)
		// var cv []byte
		// // 根据类型判断
		// if t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Uint8 {
		// 	cv = result.([]byte)
		// } else if t.Kind() == reflect.String {
		// 	cv = result.([]byte)
		// } else if t.Kind() == reflect.Interface {
		//
		// } else if t.Kind() == reflect.Struct {
		// 	cv, err = json.Marshal(result)
		// } else {
		// 	cv = result.([]byte)
		// }
		_, err = r.db.SetNX(r.ctx, key, result, expiration).Result()
	}
	return
}

// Del 删除keys
func (r *RedisEngine) Del(keys ...string) (err error) {
	err = r.db.Del(r.ctx, keys...).Err()
	return
}

// Incr key值每次加一 并返回新值
func (r *RedisEngine) Incr(key string) (result int64, err error) {
	result, err = r.db.Incr(r.ctx, key).Result()
	return
}

// IncrBy key值每次加指定数值 并返回新值
func (r *RedisEngine) IncrBy(key string, incr int64) (result int64, err error) {
	result, err = r.db.IncrBy(r.ctx, key, incr).Result()
	return
}

// IncrByFloat key值每次加指定浮点型数值 并返回新值
func (r *RedisEngine) IncrByFloat(key string, incrFloat float64) (result float64, err error) {
	result, err = r.db.IncrByFloat(r.ctx, key, incrFloat).Result()
	return
}

// Decr key值每次递减 1 并返回新值
func (r *RedisEngine) Decr(key string) (result int64, err error) {
	result, err = r.db.Decr(r.ctx, key).Result()
	return
}

// DecrBy key值每次递减指定数值 并返回新值
func (r *RedisEngine) DecrBy(key string, incr int64) (result int64, err error) {
	result, err = r.db.DecrBy(r.ctx, key, incr).Result()
	return
}

// Expire 设置 key的过期时间
func (r *RedisEngine) Expire(key string, ex time.Duration) (result bool, err error) {
	result, err = r.db.Expire(r.ctx, key, ex).Result()
	if err != nil {
		result = false
	}
	return
}

// MSet 批量设置，没有过期时间
func (r *RedisEngine) MSet(value ...interface{}) error {
	return r.db.MSet(r.ctx, value...).Err()
}

// MSet 批量获取数据
func (r *RedisEngine) MGet(keys ...string) (result []interface{}, err error) {
	result, err = r.db.MGet(r.ctx, keys...).Result()
	return
}

// Do 执行命令,返回结果
func (r *RedisEngine) Do(args ...interface{}) *redis.Cmd {
	return r.db.Do(r.ctx, args)
}

// FlushDB 清空缓存
func (r *RedisEngine) FlushDB() error {
	return r.db.FlushDB(r.ctx).Err()
}

// Publish 发布
func (r *RedisEngine) Publish(channel string, msg interface{}) error {
	return r.db.Publish(r.ctx, channel, msg).Err()
}

// Subscribe 订阅
func (r *RedisEngine) Subscribe(channel string, subscribe func(msg *redis.Message, err error)) {
	pubsub := r.db.Subscribe(r.ctx, channel)
	defer pubsub.Close()
	for {
		msg, err := pubsub.ReceiveMessage(r.ctx)
		subscribe(msg, err)
	}
}

// LPush 从列表左边插入数据，列表的头部（左边）,尾部（右边）
func (r *RedisEngine) LPush(key string, data ...interface{}) (err error) {
	err = r.db.LPush(r.ctx, key, data...).Err()
	return
}

// RPush 从列表右边边插入数据
func (r *RedisEngine) RPush(key string, data ...interface{}) (err error) {
	err = r.db.RPush(r.ctx, key, data...).Err()
	return
}

// LRange 列表从左边开始取出start至stop位置的数据
func (r *RedisEngine) LRange(key string, start, stop int64) error {
	return r.db.LRange(r.ctx, key, start, stop).Err()
}

// LPop 列表左边取出
func (r *RedisEngine) LPop(key string) *redis.StringCmd {
	return r.db.LPop(r.ctx, key)
}

// RPop 列表右边取出
func (r *RedisEngine) RPop(key string) *redis.StringCmd {
	return r.db.RPop(r.ctx, key)
}

// HSet 列表哈希插入
func (r *RedisEngine) HSet(key string, values ...interface{}) error {
	return r.db.HSet(r.ctx, key, values...).Err()
}

// HGet 列表哈希取出
func (r *RedisEngine) HGet(key, field string) *redis.StringCmd {
	return r.db.HGet(r.ctx, key, field)
}

// HMSet 列表哈希批量插入
func (r *RedisEngine) HMSet(key string, values ...interface{}) error {
	return r.db.HMSet(r.ctx, key, values...).Err()
}

// HMGet 列表哈希批量取出
func (r *RedisEngine) HMGet(key string, field ...string) []interface{} {
	return r.db.HMGet(r.ctx, key, field...).Val()
}

// SAdd 列表无序集合插入
func (r *RedisEngine) SAdd(key string, members ...interface{}) error {
	return r.db.SAdd(r.ctx, key, members...).Err()
}

// SMembers 列表无序集合，返回所有元素
func (r *RedisEngine) SMembers(key string) (result []string, err error) {
	result, err = r.db.SMembers(r.ctx, key).Result()
	return
}

// SIsMember 列表无序集合，检查元素是否存在
func (r *RedisEngine) SIsMember(key string, member interface{}) (result bool, err error) {
	result, err = r.db.SIsMember(r.ctx, key, member).Result()
	if err != nil {
		result = false
	}
	return
}

// SetMaxMemory 设置内存大小
func (r *RedisEngine) SetMaxMemory(size string) bool {
	return false
}
