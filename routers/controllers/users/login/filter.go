package login

import (
	"github.com/gin-gonic/gin"
	"net/mail"
	"regexp"
)

func checkParam(c *gin.Context, p Param) bool {
	if !isValidEmail(p.Email) {
		return false
	}
	return true
}

// isValidEmail 检查给定的电子邮件地址是否合法
func isValidEmail(email string) bool {
	// 首先使用标准库函数检查基本格式
	if _, err := mail.ParseAddress(email); err != nil {
		return false
	}

	// 使用正则表达式进行更详细的格式检查
	var emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
