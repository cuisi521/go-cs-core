// package test
// @Author cuisi
// @Date 2023/12/26 17:59:00
// @Desc
package jwt

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/cuisi521/go-cs-core/store/cache"
)

type JwtCache struct {
	cache cache.Cacher
}

func InstallCache() *JwtCache {
	c := cache.NewMemCache()
	c.SetMaxMemory("10MB")
	return &JwtCache{cache: c}
}

func (j *JwtCache) CreateToken(key []byte, pl *StandardClaims) (string, *JwtResult) {
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

func (j *JwtCache) ParseToken(token string, key []byte) (*StandardClaims, *JwtResult) {
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

	// 使用key将token中的头部和负载用.连接后进行签名
	// 这个签名应该个token中第三部分的签名一致
	// confirmSignature, err := generateSignature(key, []byte(encodedHeader+"."+encodedPayload))
	// if err != nil {
	// 	return nil, Result("生成签名错误", SignaErr)
	// }
	// // 如果不一致
	// if signature != confirmSignature {
	// 	return nil, Result("token验证失败", Fail)
	// }

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
			if token != r.(string) {
				return nil, Result(err.Error(), Fail)
			}
		}
	}
	// 返回我们的JWT对象以供后续使用
	// return &Jwt{encodedHeader, string(dstPayload), signature}, nil
	return sc, nil
}
