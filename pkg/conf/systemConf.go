package conf

import (
	"GoCloud/pkg/log"
	"sync"
)

// system 系统通用配置
type system struct {
	Mode  string `map`
	Debug bool
	once  sync.Once
	//Listen        string
	SessionSecret string
	HashIDSalt    string
	//GracePeriod   int
	//ProxyHeader   string
}

var systemConfig = new(system)

// SystemConfig 系统公用静态配置
func SystemConfig() *system {
	systemConfig.once.Do(func() {
		log.NewEntry("init systemConfig...")
		Config().Sub("system").Unmarshal(&systemConfig)
		log.NewEntry("init systemConfig...end")
	})
	return systemConfig
}
