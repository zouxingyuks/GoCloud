package cache

import (
	"github.com/pkg/errors"
	"time"
)

type Second int
type Config any

// Driver 键值缓存存储容器接口
type Driver interface {
	// Set 设置值，ttl为过期时间，单位为秒
	Set(key string, value any, ttl Second) error

	// Get 取值，并返回是否成功
	Get(key string) (any, error)

	// Gets 批量取值，返回成功取值的map以及不存在的值的列表
	Gets(keys []string, prefix string) (map[string]interface{}, []string, error)

	// Sets 批量设置值，所有的key都会加上prefix前缀
	Sets(values map[string]interface{}, prefix string) error

	/* H 结构化存储 */

	// HGet 获取值的某个字段
	HGet(key string, field string) (any, error)

	// HMGet 获取值的多个字段
	HMGet(key string, fields []string) (map[string]interface{}, error)

	// HGetAll 获取值的所有字段
	HGetAll(key string) (map[string]string, error)

	// HMSet 设置值的多个字段
	HMSet(key string, values map[string]interface{}, ttl Second) error

	// HSet 设置值的某个字段
	HSet(key string, field string, value interface{}, ttl Second) error

	// Delete 删除值
	Delete(keys []string, prefix string) error

	// Persist 将缓存数据保存到磁盘
	Persist(path string) error

	// Restore 从磁盘恢复缓存数据
	Restore(path string) error
}

type driver struct {
	// 缓存容器
	container Driver
	// 容器配置
	config any
}
type Option interface {
	apply(*driver)
}

type function struct {
	f func(*driver)
}

func (f0 function) apply(o *driver) {
	f0.f(o)
}
func newFunction(f func(option *driver)) Option {
	return &function{
		f: f,
	}

}

func New(k Kind, options ...Option) (Driver, error) {
	d := &driver{}
	for _, o := range options {
		o.apply(d)
	}
	//todo 选择容器
	switch k {
	case KindMemory:
		panic("not implement")
		//d.container = newMemoryStore()
	case KindRedis:
		if d.config == nil {
			return nil, errors.New("cache redis need config")
		}
		d.container = newRedisStore(d.config.(RedisConfig))
	}
	return d.container, nil
}

func WithConf[T RedisConfig | TempConfig](conf T) Option {
	return newFunction(func(o *driver) {
		o.config = conf
	})
}

type TempConfig struct {
}

func makeTime(ttl Second) time.Duration {
	//1. 如果ttl为-1则代表永不过期
	if ttl == -1 {
		return -1
	}
	//2. 如果ttl为0则代表使用默认过期时间:10分钟
	if ttl == 0 {
		return time.Duration(600) * time.Second
	}

	return time.Duration(ttl) * time.Second
}
