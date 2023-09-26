package conf

import "sync"

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
	once        sync.Once
}

var databaseConfig = new(database)

// DatabaseConfig 数据库配置
func DatabaseConfig() *database {
	databaseConfig.once.Do(func() {
		Config().Sub("database").Unmarshal(&databaseConfig)
	})
	return databaseConfig

}
