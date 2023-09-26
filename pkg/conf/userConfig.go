package conf

import "sync"

type user struct {
	DefaultGroup int
	once         sync.Once
}

var userConfig = new(user)

// UserConfig 系统公用静态配置
func UserConfig() *user {
	userConfig.once.Do(func() {
		Config().Unmarshal(&userConfig)
	})
	return userConfig
}
