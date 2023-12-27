// package jwt
// @Author cuisi
// @Date 2023/12/25 17:10:00
// @Desc
package jwt

type StandardClaims struct {
	UseKey    string `json:"uk,omitempty"`
	IssuedAt  int64  `json:"iat,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	Payload   any    `json:"pld,omitempty"`
}

type JwtResult struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func Result(message string, code int) *JwtResult {
	return &JwtResult{Message: message, Code: code}
}
