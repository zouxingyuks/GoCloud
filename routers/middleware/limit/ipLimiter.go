package limit

import (
	"GoCloud/pkg/log"
	"GoCloud/service/serializer"
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
	"net/http"
	"sync"
)

// todo 限流器的释放，ip限流器在长期运行的情况下会出现内存泄漏
type ipLimiter struct {
	limiter *rate.Limiter
}

var (
	ips = make(map[string]*ipLimiter)
	mu  sync.RWMutex
)

// getLimiter 获取ip对应的限流器
func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()
	// 从map中获取ip对应的限流器
	lim, exists := ips[ip]
	if !exists {
		lim = &ipLimiter{limiter: rate.NewLimiter(1, 3)}
		ips[ip] = lim
	}
	return ips[ip].limiter
}

// WithIPLimiter  ip限流器
func WithIPLimiter() Option {
	log.NewEntry("middleware.limit").Info("load ip limiter middleware")
	return newFuncOption(func(o *option) {
		entry := log.NewEntry("middleware.limit.ipLimiter")

		o.fs = append(o.fs, func(c *gin.Context) {
			ip := c.ClientIP()
			limiter := getLimiter(ip)
			entry.Debug("ip limiter", log.Field{Key: "ip", Value: ip})
			if !limiter.Allow() {
				res := serializer.NewResponse(entry, http.StatusTooManyRequests, serializer.WithMsg("Too Many Requests"))
				c.AbortWithStatusJSON(res.Code, res)
				return
			}
		})
	})
}
