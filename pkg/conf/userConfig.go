package conf

import (
	"log"
	"sync"
)

// todo 默认群组的管理
type userController struct {
	DefaultGroup int
	EmailVerify  bool
	once         sync.Once
}

var userConfig = new(userController)

// UserConfig 系统公用静态配置
func UserConfig() *userController {
	userConfig.once.Do(func() {
		log.Println("init UserConfig...")
		Config().Sub("usercontroller").Unmarshal(&userConfig)
		log.Println("init UserConfig success")
	})
	return userConfig
}
