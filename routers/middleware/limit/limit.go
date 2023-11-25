package limit

import (
	"GoCloud/pkg/cache"
	"GoCloud/pkg/log"
	"github.com/gin-gonic/gin"
	"sync"
)

// 类型名称及字段的首字母小写（包内私有）
type option struct {
	fs []func(c *gin.Context)
}
type Option interface {
	apply(*option)
}

// funcOption 定义funcOption类型，实现 IOption 接口
type funcOption struct {
	f func(*option)
}

func (fo funcOption) apply(o *option) {
	fo.f(o)
}

func newFuncOption(f func(option *option)) Option {
	return &funcOption{
		f: f,
	}
}

// 使用中间件进行访问控制

var cacheInstance = new(struct {
	cache.Driver
	sync.Once
})

// New 接口调用频率限制
func New(opts ...Option) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.NewEntry("middleware.limit").Debug("load limit middleware")
		o := new(option)
		for _, opt := range opts {
			opt.apply(o)
		}
		for _, f := range o.fs {
			f(c)
		}
		c.Next()
	}
}
