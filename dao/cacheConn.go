package dao

import (
	"GoCloud/conf"
	"GoCloud/pkg/cache"
	"GoCloud/pkg/log"
	"sync"
)

var cacheInstance = new(struct {
	cache.Driver
	sync.Once
})

func Cache() cache.Driver {
	cacheInstance.Do(
		func() {
			var err error
			switch conf.DaoConfig().Cache.Kind {
			case cache.KindRedis:
				cacheInstance.Driver, err = cache.New(cache.KindRedis, cache.WithConf(cache.RedisConfig{
					Addr:     conf.DaoConfig().Cache.Redis.Server,
					Password: conf.DaoConfig().Cache.Redis.Password,
					DB:       conf.DaoConfig().Cache.Redis.DB,
					PoolSize: conf.DaoConfig().Cache.Redis.PoolSize,
				}))
				if err != nil {
					log.NewEntry("service.user").Error("cache init error", log.Field{
						Key:   "err",
						Value: err,
					})
					log.NewEntry("service.user").Info("cache use memory")
					cacheInstance.Driver, err = cache.New(cache.KindMemory)
				}

			default:
				cacheInstance.Driver, err = cache.New(cache.KindMemory)
			}
			if err != nil {
				log.NewEntry("service.user").Error("cache init error", log.Field{
					Key:   "err",
					Value: err,
				})
			}
		})

	return cacheInstance.Driver

}
