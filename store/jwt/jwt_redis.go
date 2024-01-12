// package jwt
// @Author cuisi
// @Date 2023/12/26 17:59:00
// @Desc
package jwt

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/cuisi521/go-cs-core/store/cache"
	"github.com/cuisi521/go-cs-core/sys/ccfg"
)

type JwtRedis struct {
	cache *cache.RedisEngine
}

func InstallRedis() *JwtRedis {
	alias := ccfg.SysCnf().Token.Redis
	if cache.Redis(alias) == nil {
		cnf := ccfg.SysCnf().Cache.Redis[alias]
		redisConfig := &cache.Config{
			Address:       cnf.Host,
			Db:            cnf.Db,
			User:          "",
			Pass:          cnf.Password,
			ReadTimeout:   time.Second * 3,
			WriteTimeout:  time.Second * 3,
			WaitTimeout:   time.Second * 10,
			TLS:           false,
			TLSSkipVerify: false,
			TLSConfig:     nil,
			SlaveOnly:     false,
			PoolSize:      cnf.PoolSize,
		}
		redisCnf := &cache.RedisCnfs{Alias: alias, Cnf: redisConfig}
		redisCnfs := []*cache.RedisCnfs{}
		redisCnfs = append(redisCnfs, redisCnf)
		err := cache.New(redisCnfs)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return &JwtRedis{cache: cache.Redis(alias)}
}

func (j *JwtRedis) CreateToken(key []byte, pl *StandardClaims) (string, *JwtResult) {
	if pl.UseKey == "" {
		pl.UseKey = uuid.New().String()
	}
	pl.IssuedAt = time.Now().Unix()

	// 将负载的数据转换为json
	payload, jsonErr := json.Marshal(pl)
	if jsonErr != nil {
		return "", Result("负载json解析错误", ParamErr)
	}

	// 将头部和负载通过base64编码，并使用.作为分隔进行连接
	encodedHeader := encodeBase64(Header)
	encodedPayload := encodeBase64(string(payload))
	HeaderAndPayload := encodedHeader + "." + encodedPayload

	// 使用签名使用的key将传入的头部和负载连接所得的数据进行签名
	signature, err := generateSignature(key, []byte(HeaderAndPayload))
	if err != nil {
		return "", Result(err.Error(), SignaErr)
	}
	dbStr := HeaderAndPayload + "." + signature
	if pl.ExpiresAt <= 0 {
		j.cache.Set(pl.UseKey, dbStr)
	} else {
		j.cache.SetEX(pl.UseKey, dbStr, time.Second*time.Duration(pl.ExpiresAt-pl.IssuedAt))
	}

	// 将token的三个部分使用.进行连接并返回
	return HeaderAndPayload + "." + signature, nil
}

func (j *JwtRedis) ParseToken(token string, key []byte) (*StandardClaims, *JwtResult) {
	// 分解规定，我们使用.进行分隔，所以我们通过.进行分隔成三个字符串的数组
	jwtParts := strings.Split(token, ".")
	// 数据数组长度不是3就说明token在格式上就不合法
	if len(jwtParts) != 3 {
		return nil, Result("无效taoke", Invalid)
	}

	// 分别拿出
	// encodedHeader := jwtParts[0]
	encodedPayload := jwtParts[1]
	// signature := jwtParts[2]

	// 将payload解base64编码
	dstPayload, _ := base64.RawURLEncoding.DecodeString(encodedPayload)
	var sc *StandardClaims
	err := json.Unmarshal(dstPayload, &sc)
	if err != nil || sc == nil || sc.UseKey == "" {
		return nil, Result("token验证失败", Fail)
	} else {
		if sc != nil {
			r, err := j.cache.Get(sc.UseKey)
			if err != nil {
				return nil, Result(err.Error(), Fail)
			}
			if token != r {
				return nil, Result(err.Error(), Fail)
			}
		}
	}
	// 返回我们的JWT对象以供后续使用
	// return &Jwt{encodedHeader, string(dstPayload), signature}, nil
	return sc, nil
}
