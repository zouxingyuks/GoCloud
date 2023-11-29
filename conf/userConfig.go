package conf

type userService struct {
	RegisterEnable bool
	EmailVerify    bool

	// 登录频率限制
	LoginLimit struct {
		// 登录频率限制开关
		Enable bool
		// 登录频率限制时间
		Period int
		// 登录频率限制次数
		Count int
	}
}
