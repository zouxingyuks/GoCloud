package middleware

import (
	"GoCloud/conf"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var secret = []byte(conf.StaticConfig().Session.Secret)

// Session 初始化session
func Session() gin.HandlerFunc {
	var err error
	store := sessions.Store(nil)
	switch conf.StaticConfig().Session.Store {
	case "redis":
		store, err = redis.NewStore(conf.RedisConfig().PoolSize, conf.RedisConfig().Network, conf.RedisConfig().Server, conf.RedisConfig().Password, secret)
		if err != nil {
			panic(err)
		}
	case "cookie":
		store = cookie.NewStore(secret)

	case "memory":
		store = memstore.NewStore(secret)

	}

	sameSiteMode := http.SameSiteDefaultMode
	switch strings.ToLower(conf.CORSConfig().SameSite) {
	case "default":
		sameSiteMode = http.SameSiteDefaultMode
	case "none":
		sameSiteMode = http.SameSiteNoneMode
	case "strict":
		sameSiteMode = http.SameSiteStrictMode
	case "lax":
		sameSiteMode = http.SameSiteLaxMode
	default:
		sameSiteMode = http.SameSiteDefaultMode
	}

	//// Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{
		HttpOnly: conf.SystemConfig().Sessions.Option.HttpOnly,
		//86400 是一天
		MaxAge:   conf.SystemConfig().Sessions.Option.MaxAge,
		Path:     "/",
		SameSite: sameSiteMode,
		Secure:   conf.SystemConfig().Sessions.Option.Secure,
	})

	return sessions.Sessions("gocloud-session", store)
}
