package routers

import (
	"GoCloud/routers/controllers"
	"github.com/gin-gonic/gin"
)

// @title GoCloud
// @version 0.0
// @description 这是一个用 go 语言实现的网盘
// @termsOfService http://swagger.io/terms/

// @contact.name 这里写联系人信息
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 这里写接口服务的host
// @BasePath /api/v0

type apiV0 struct {
}

// 遵循 restful api 设计原则
// 使用名词复数来表示资源集合
func (apiV0) load(r *gin.Engine) {
	v0 := r.Group("/api/v0")
	/*
		中间件
	*/
	//apiV0.Use(middleware.Session(conf.SystemConfig.SessionSecret))

	//跨域相关
	//InitCORS(r)

	//// 测试模式加入Mock助手中间件
	//if gin.Mode() == gin.TestMode {
	//	v3.Use(middleware.MockHelper())
	//}
	//// 用户会话
	//v3.Use(middleware.CurrentUser())
	//
	//// 禁止缓存
	//v3.Use(middleware.CacheControl())
	//
	/*
		路由
	*/
	{
		// 用户相关路由
		user := v0.Group("users")
		{
			//用户登录
			user.POST("session", middleware.CaptchaRequired("login_captcha"), controllers.V0{}.UserLogin)
			// 用户注册
			user.POST("",
				//todo 允许设置注册
				//middleware.IsFunctionEnabled("register_enabled"),
				//middleware.CaptchaRequired("reg_captcha"),
				controllers.V0{}.UserRegister,
			)
			//// 用二步验证户登录
			//user.POST("2fa", controllers.User2FALogin)
			//// 发送密码重设邮件
			//user.POST("reset", middleware.CaptchaRequired("forget_captcha"), controllers.UserSendReset)
			//// 通过邮件里的链接重设密码
			//user.PATCH("reset", controllers.UserReset)
			//// 邮件激活
			//user.GET("activate/:id",
			//	middleware.SignRequired(auth.General),
			//	middleware.HashID(hashid.UserID),
			//	controllers.UserActivate,
			//)
			//// WebAuthn登陆初始化
			//user.GET("authn/:username",
			//	middleware.IsFunctionEnabled("authn_enabled"),
			//	controllers.StartLoginAuthn,
			//)
			//// WebAuthn登陆
			//user.POST("authn/finish/:username",
			//	middleware.IsFunctionEnabled("authn_enabled"),
			//	controllers.FinishLoginAuthn,
			//)
			//// 获取用户主页展示用分享
			//user.GET("profile/:id",
			//	middleware.HashID(hashid.UserID),
			//	controllers.GetUserShare,
			//)
			//// 获取用户头像
			//user.GET("avatar/:id/:size",
			//	middleware.HashID(hashid.UserID),
			//	middleware.StaticResourceCache(),
			//	controllers.GetUserAvatar,
			//)
		}

	}

}
