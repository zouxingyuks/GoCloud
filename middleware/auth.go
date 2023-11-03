package middleware

import (
	"GoCloud/pkg/log"
	"GoCloud/pkg/util"
	"github.com/gin-gonic/gin"
)

// Captcha 验证请求签名
func Captcha(configName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// todo 验证码
		//if !captcha.VerifyString(captchaId, captchaSolution) {
		//	c.JSON(400, serializer.Err(
		//		serializer.CodeCaptchaErr,
		//		"验证码错误",
		//		nil))
		//	c.Abort()
		//	return
		//}
		//c.Next()
	}
}

//

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid := util.Session().Get(c, "user_id")
		//if uid != nil {
		//	user, err := dao.GetUser(
		//		dao.WithStatus(dao.UserActive),
		//		dao.WithUUID(uid.(string)),
		//	)
		//	if err == nil {
		//		c.Set("user", &user)
		//	} else {
		//		log.NewEntry("middleware.current_user").Error("Failed to get user: ", log.Field{
		//			Key:   "err",
		//			Value: err,
		//		})
		//	}
		//}
		if uid == nil {
			c.JSON(401, gin.H{
				"msg": "未登录",
			})
			c.Abort()
		}
		log.NewEntry("middleware.current_user").Debug("Get current user: ", log.Field{
			Key:   "uuid",
			Value: uid,
		})
		c.Next()
	}
}
