package conf

import (
	"GoCloud/pkg/log"
	"github.com/pkg/errors"
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
}

var corsConfig = new(struct {
	once sync.Once
	*cors
})

// CORSConfig 跨域配置
func CORSConfig() *cors {
	corsConfig.once.Do(func() {
		log.NewEntry("conf").Info("init CORSConfig...start")
		err := Config().Sub("cors").Unmarshal(&corsConfig.cors)
		if err != nil {
			log.NewEntry("conf").Error(errors.Wrap(err, "init CORSConfig error").Error())
			log.NewEntry("conf").Info("use default CORSConfig")
			temp, _ := defaultConfig["cors"].(cors)
			corsConfig.cors = &temp
		}

		log.NewEntry("conf").Info("init CORSConfig...end")
	})
	return corsConfig.cors
}
