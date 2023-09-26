package conf

import "sync"

// redis 配置
// redis 可以使用动态配置
type redis struct {
	Network  string
	Server   string
	User     string
	Password string
	DB       string
	once     sync.Once
}

var redisConfig = new(redis)

// RedisConfig Redis服务器配置
func RedisConfig() *redis {
	redisConfig.once.Do(
		func() {
			Config().Sub("redis").Unmarshal(&redisConfig)
		})
	return redisConfig
}
