package conf

import (
	"GoCloud/pkg/log"
	"sync"
)

type userController struct {
	EmailVerify bool
	// 登录频率限制
	LoginLimit struct {
		// 登录频率限制开关
		Enable bool
		// 登录频率限制时间
		Period int
		// 登录频率限制次数
		Count int
	}
}

var userConfig = new(struct {
	once sync.Once
	*userController
})

// UserConfig 系统公用静态配置
func UserConfig() *userController {
	userConfig.once.Do(func() {
		entry := log.NewEntry("conf.userConfig")
		entry.Debug("init UserConfig")
		Config().Sub("usercontroller").Unmarshal(&userConfig.userController)
		entry.Debug("init UserConfig", log.Field{
			Key:   "userConfig",
			Value: userConfig,
		})
		entry.Info("init UserConfig end")
	})
	return userConfig.userController
}
