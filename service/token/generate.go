package token

import (
	"GoCloud/conf"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

func Generate(action, sub string, exp time.Duration) (string, error) {
	claims := &struct {
		jwt.RegisteredClaims
		Action string `json:"act"`
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   sub,                                     // 用户ID
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),          // 签发时间

		},
		Action: action,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.StaticConfig().Session.Secret))
}
