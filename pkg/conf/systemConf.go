package conf

// system 系统通用配置
type system struct {
	Mode string
	//Listen        string
	Debug bool
	//SessionSecret string
	//HashIDSalt    string
	//GracePeriod   int
	//ProxyHeader   string
}

const (
	/*
		SystemConfig 系统配置部分
	*/
	/*
		UserConfig 用户管理相关配置
	*/
	UserConfigDefaultGroup = "user.defaultGroup"
)

// SystemConfig 系统公用静态配置
var SystemConfig = &system{
	Debug: Config().GetBool("SystemConfig.Debug"),
	Mode:  Config().GetString("SystemConfig.Mode"),
}

// Get 用于获取动态信息
// todo 是否可能出现安全漏洞
func (system) Get(attr string) any {
	switch attr {
	case UserConfigDefaultGroup:
		//用户群组以数字表示
		return Config().GetInt(UserConfigDefaultGroup)
	default:
		return nil
	}
}

func (s system) Set(attr string, value any) error {
	switch attr {
	case UserConfigDefaultGroup:
		//todo 在 viper 中设置默认群组
	}
	//TODO implement me
	panic("implement me")
}

func (s system) Save() error {
	//TODO implement me
	panic("implement me")
}
