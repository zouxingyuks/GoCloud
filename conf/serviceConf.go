package conf

import (
	"GoCloud/pkg/log"
	"sync"
)

type service struct {
	//用户服务
	User userService
}

var serviceConfig = new(struct {
	once sync.Once
	*service
})

// ServiceConfig 服务配置
func ServiceConfig() *service {
	serviceConfig.once.Do(func() {
		entry := log.NewEntry("conf.serviceConfig")
		entry.Debug("init ServiceConfig")
		err := Config().Sub("service").Unmarshal(&serviceConfig.service)
		if err != nil {
			entry.Panic("init ServiceConfig error", log.Field{
				Key:   "err",
				Value: err,
			})
		}
		entry.Debug("init ServiceConfig end")
		entry.Info("init ServiceConfig success", log.Field{
			Key:   "ServiceConfig",
			Value: serviceConfig.service,
		})
	})
	return serviceConfig.service
}
