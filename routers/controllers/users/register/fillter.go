package register

import (
	"GoCloud/pkg/log"
	"github.com/gin-gonic/gin"
	"net/mail"
	"regexp"
	"unicode"
)

func checkParam(c *gin.Context, p Param) bool {
	if !isValidEmail(p.Email) || !isValidPassword(p.Password) {
		entry := log.NewEntry("controller.user.login")
		// 记录日志
		entry.Warn(p.Email+"尝试登陆，但是参数错误", log.Field{
			Key:   "ip",
			Value: c.ClientIP(),
		})
		//此处不记录密码，主要是为了防止日志泄露，以及逻辑上不应该记录密码
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

// isValidPassword 检查给定的密码是否合法
func isValidPassword(password string) bool {
	var (
		hasMinLength = len(password) >= 8
		hasMaxLength = len(password) <= 20
		hasUpper     bool
		hasLower     bool
		hasNumber    bool
		hasSpecial   bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasMinLength && hasMaxLength && hasUpper && hasLower && hasNumber && hasSpecial
}
