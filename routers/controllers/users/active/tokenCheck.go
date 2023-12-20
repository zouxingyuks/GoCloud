package active

import (
	"GoCloud/conf"

	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

func Parse(tokenStr, act string) (any, error) {
	// 2.解析token
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.StaticConfig().Session.Secret), nil
	})
	if err != nil {
		return nil, errors.Wrap(err, "jwt.Parse failed")
	}

	/* 校验token
	校验token是否有效
	校验token中的 sub 是否有效
	校验token中的act是否匹配
	*/
	if claims, ok := token.Claims.(jwt.MapClaims); token.Valid && ok && claims["sub"] != nil && claims["act"] == act {
		sub := claims["sub"]
		return sub, nil
	}
	return nil, errors.Wrap(err, "invalided token")
}
