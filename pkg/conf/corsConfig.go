package conf

import (
	"log"
	"sync"
)

// cors 跨域配置
type cors struct {
	AllowOrigins     []string
	AllowMethods     []string
	AllowHeaders     []string
	AllowCredentials bool
	ExposeHeaders    []string
	SameSite         string
	Secure           bool
	once             sync.Once
}

var corsConfig = new(cors)

// CORSConfig 跨域配置
func CORSConfig() *cors {
	corsConfig.once.Do(func() {
		log.Println("init CORSConfig...")
		Config().Sub("cors").Unmarshal(&corsConfig)
		log.Println("init CORSConfig...end")
	})
	return corsConfig
}
