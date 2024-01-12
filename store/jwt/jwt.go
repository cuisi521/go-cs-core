// package jwt
// @Author cuisi
// @Date 2023/12/25 14:59:00
// @Desc
package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"
)

const (
	// 成功
	Success = 0
	// 失败
	Fail = 1
	// 参数错误
	ParamErr = 2
	// 签名错误
	SignaErr = 3
	// 无效token
	Invalid = 4
	// 过期
	PastDue = 5
	// 标准头部
	Header = `{"alg":"HS256","typ":"JWT"}`
)

var (
	jwter Jwter
)

type Jwter interface {
	CreateToken(key []byte, pl *StandardClaims) (string, *JwtResult)
	ParseToken(token string, key []byte) (*StandardClaims, *JwtResult)
}

type Jwt struct {
	header         string `json:"header"`
	StandardClaims string `json:"standardClaims"`
	signature      string `json:"signature"`
}

func RegisterCache() {
	jwter = InstallJwt()
	// if ccfg.SysCnf() == nil {
	// 	jwter = InstallJwt()
	// 	return
	// }
	// switch ccfg.SysCnf().Token.Mod {
	// case 0:
	// 	jwter = InstallJwt()
	// case 1:
	// 	jwter = InstallCache()
	// case 2:
	// 	jwter = InstallRedis()
	// default:
	// 	jwter = InstallJwt()
	// }
}

func InstallJwt() *Jwt {
	return &Jwt{}
}

func encodeBase64(data string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(data))
}

// generateSignature 生成签名
// @author By Cuisi 2023/12/25 15:01:00
func generateSignature(key []byte, data []byte) (string, error) {
	// 创建一个哈希对象
	hash := hmac.New(sha256.New, key)
	// 将要签名的信息写入哈希对象中 hash.Write(data)
	_, err := hash.Write(data)
	if err != nil {
		return "", err
	}
	// hash.Sum()计算签名，在这里会返回签名内容
	// 将签名经过base64编码生成字符串形式返回。
	return encodeBase64(string(hash.Sum(nil))), nil
}

// CreateToken 创建token
// @author By Cuisi 2023/12/25 15:01:00
func (j *Jwt) CreateToken(key []byte, pl *StandardClaims) (string, *JwtResult) {
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

	// 将token的三个部分使用.进行连接并返回
	return HeaderAndPayload + "." + signature, nil
}

// ParseJwt 解析token
// @author By Cuisi 2023/12/25 15:02:00
func (j *Jwt) ParseToken(token string, key []byte) (*StandardClaims, *JwtResult) {
	// 分解规定，我们使用.进行分隔，所以我们通过.进行分隔成三个字符串的数组
	jwtParts := strings.Split(token, ".")
	// 数据数组长度不是3就说明token在格式上就不合法
	if len(jwtParts) != 3 {
		return nil, Result("无效taoke", Invalid)
	}

	// 分别拿出
	encodedHeader := jwtParts[0]
	encodedPayload := jwtParts[1]
	signature := jwtParts[2]

	// 使用key将token中的头部和负载用.连接后进行签名
	// 这个签名应该个token中第三部分的签名一致
	confirmSignature, err := generateSignature(key, []byte(encodedHeader+"."+encodedPayload))
	if err != nil {
		return nil, Result("生成签名错误", SignaErr)
	}
	// 如果不一致
	if signature != confirmSignature {
		return nil, Result("token验证失败", Fail)
	}

	// 将payload解base64编码
	dstPayload, _ := base64.RawURLEncoding.DecodeString(encodedPayload)
	var sc *StandardClaims
	err = json.Unmarshal(dstPayload, &sc)
	if err != nil {
		return nil, Result("token验证失败", Fail)
	} else {
		if sc != nil && sc.ExpiresAt < time.Now().Unix() {
			return nil, Result("token已过期", PastDue)
		}
	}
	// 返回我们的JWT对象以供后续使用
	// return &Jwt{encodedHeader, string(dstPayload), signature}, nil
	return sc, nil
}

func CreateToken(key []byte, pl *StandardClaims) (string, *JwtResult) {
	return jwter.CreateToken(key, pl)
}

func ParseToken(token string, key []byte) (*StandardClaims, *JwtResult) {
	return jwter.ParseToken(token, key)
}
