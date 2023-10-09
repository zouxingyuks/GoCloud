package hash

import "sync"

// iHash 是用于数据混淆的接口
type iHash interface {
	// Encode 计算哈希值
	Encode(v []int) (string, error)
	// Decode  对给定数据计算原始数据
	Decode(raw string) ([]int, error)
}

var hashInstance = new(struct {
	iHash
	sync.Once
})

func Encode(v []int) (string, error) {
	hashInstance.Once.Do(func() {
		hashInstance.iHash = newHash()
	})
	return hashInstance.Encode(v)
}
func Decode(raw string) ([]int, error) {
	hashInstance.Once.Do(func() {
		hashInstance.iHash = newHash()
	})
	return hashInstance.Decode(raw)
}
