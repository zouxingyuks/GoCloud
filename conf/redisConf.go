package conf

import (
	"GoCloud/pkg/log"
	"sync"
)

// redis 配置
// redis 可以使用动态配置
type redis struct {
	Network  string
	Server   string
	User     string
	Password string
	PoolSize int
	DB       int
}

var redisConfig = new(struct {
	once sync.Once
	*redis
})

// RedisConfig Redis服务器配置
func RedisConfig() *redis {
	redisConfig.once.Do(
		func() {
			log.NewEntry("conf").Info("init redisConfig...start")
			err := Config().Sub("redis").Unmarshal(&redisConfig.redis)
			if err != nil {
				log.NewEntry("conf").Warn("init redisConfig...failed", log.Field{
					Key:   "error",
					Value: err,
				})
			} else {
				log.NewEntry("conf").Info("init redisConfig success")
			}
		})
	return redisConfig.redis
}
