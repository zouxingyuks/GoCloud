package conf

// redis 配置
// redis 可以使用动态配置
type redis struct {
	Network  string
	Server   string
	User     string
	Password string
	DB       string
}

// RedisConfig Redis服务器配置
var RedisConfig = &redis{
	Network:  "tcp",
	Server:   "",
	Password: "",
	DB:       "0",
}
