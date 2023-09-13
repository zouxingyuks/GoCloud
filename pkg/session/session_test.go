package session

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAndSetSession(t *testing.T) {
	r := gin.Default()
	// 配置session中间件
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/test", func(c *gin.Context) {
		session := NewSession(zaptest.NewLogger(t))
		session.Set(c, map[string]interface{}{
			"num": "111",
		})
		value := session.Get(c, "num")
		if value != "111" {
			t.Errorf("Expected value '111', got %v", value)
		}
		fmt.Println("value:", value)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	expectedStatus := http.StatusOK
	assert.Equal(t, expectedStatus, w.Code)

	fmt.Printf("Body: %+v", w.Body)
}

func TestDeleteSession(t *testing.T) {
	r := gin.Default()
	// 配置session中间件
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/test", func(c *gin.Context) {
		session := NewSession(zaptest.NewLogger(t))
		session.Set(c, map[string]interface{}{
			"num": "111",
		})
		value := session.Get(c, "num")
		if value != "111" {
			t.Errorf("Expected value '111', got %v", value)
		}
		fmt.Println("value:", value)
		session.Delete(c, "num")
		value = session.Get(c, "num")
		if value != nil {
			t.Errorf("Expected value 'nil', got %v", value)
		}
		fmt.Println("value:", value)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	expectedStatus := http.StatusOK
	assert.Equal(t, expectedStatus, w.Code)

	fmt.Printf("Body: %+v", w.Body)
}
