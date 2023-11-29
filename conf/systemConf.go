package conf

import (
	"GoCloud/pkg/log"
	"github.com/pkg/errors"
	"sync"
)

// system 系统通用配置
type system struct {
	Debug      bool
	HashIDSalt string
}

var systemConfig = new(struct {
	once sync.Once
	*system
})

// SystemConfig 系统公用静态配置
func SystemConfig() *system {
	systemConfig.once.Do(func() {
		systemConfig.system = new(system)
		log.NewEntry("conf").Info("init systemConfig...start")
		err := Config().Sub("system").Unmarshal(&systemConfig.system)
		if err != nil {
			log.NewEntry("conf").Panic(errors.New("init systemConfig...failed").Error())
		}
		log.SetDebug(systemConfig.Debug)
		log.NewEntry("conf").Info("init systemConfig...done")
	})
	return systemConfig.system
}
