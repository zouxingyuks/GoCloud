package cache

import "sync"

type Second int

var driverInstance = new(struct {
	iDriver
	once sync.Once
})

// iDriver 键值缓存存储容器
type iDriver interface {
	// Set 设置值，ttl为过期时间，单位为秒
	Set(key string, value interface{}, ttl Second) error

	// Get 取值，并返回是否成功
	Get(key string) (interface{}, error)

	// Gets 批量取值，返回成功取值的map即不存在的值
	Gets(keys []string, prefix string) (map[string]interface{}, []string, error)

	// Sets 批量设置值，所有的key都会加上prefix前缀
	Sets(values map[string]interface{}, prefix string) error

	// Delete 删除值
	Delete(keys []string, prefix string) error

	// Persist Save in-memory driverInstance to disk
	Persist(path string) error

	// Restore driverInstance from disk
	Restore(path string) error
}

func Driver() iDriver {
	driverInstance.once.Do(
		func() {
			//todo 为测试模式编写
			//driverInstance=NewMemoStore()
			driverInstance.iDriver = RedisStore()
		})

	return driverInstance.iDriver
}
