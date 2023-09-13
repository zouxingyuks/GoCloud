package session

import (
	sessions "github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var log *zap.Logger

type Session interface {
	Get(c *gin.Context, key string) interface{}
	Set(c *gin.Context, list map[string]interface{})
	Delete(c *gin.Context, key string)
}

func NewSession(l *zap.Logger) Session {
	log = l
	return &session{}
}

type session struct {
}

// Get 获取session
func (session) Get(c *gin.Context, key string) interface{} {
	s := sessions.Default(c)
	return s.Get(key)
}

// Set 设置session
func (session) Set(c *gin.Context, list map[string]interface{}) {
	s := sessions.Default(c)
	for key, value := range list {
		s.Set(key, value)
	}
	err := s.Save()
	if err != nil {
		log.Warn("无法设置 Session 值...",
			zap.Error(err),
		)
	}
}

// Delete 删除session
func (session) Delete(c *gin.Context, key string) {
	s := sessions.Default(c)
	s.Delete(key)
	s.Save()
}
