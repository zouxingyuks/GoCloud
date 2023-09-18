package conf

type SystemConfig interface {
}

const (
	/*
		SystemConfig 系统配置部分
	*/
	SystemConfigDebug = "Debug"
	SystemConfigMode  = "Mode"

	/*
		UserConfig 用户管理相关配置
	*/
	UserConfigDefaultGroup = "user.defaultGroup"
)

// systemConfig 系统公用静态配置
var systemConfig = &system{
	Debug: false,
	Mode:  "master",
}

// 为了实现动态与静态的动态修改，此处改成使用函数来获取信息
func GetSystemConfig(attr string) any {
	switch attr {
	case SystemConfigDebug:
		return systemConfig.Debug
	case SystemConfigMode:
		return systemConfig.Mode
	case UserConfigDefaultGroup:
		//todo  从 viper 获取默认群组
		return 1
	default:
		return nil
	}
}
