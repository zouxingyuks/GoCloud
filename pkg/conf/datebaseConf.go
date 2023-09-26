package conf

// database 数据库
// 数据库配置不需要动态加载
type database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	DBFile      string
	Port        int
	Charset     string
	UnixSocket  bool
}

// DatabaseConfig 数据库配置
var DatabaseConfig = &database{
	Type:       Config().GetString("DatabaseConfig.Type"),
	Charset:    Config().GetString("DatabaseConfig.Charset"),
	DBFile:     Config().GetString("DatabaseConfig.DBFile"),
	Port:       Config().GetInt("DatabaseConfig.Port"),
	UnixSocket: Config().GetBool("DatabaseConfig.UnixSocket"),
}
