package conf

import (
	"GoCloud/pkg/log"
	"sync"
)

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
type dao struct {
	Database database
}

var daoConfig = new(struct {
	*dao
	sync.Once
})

func DaoConfig() *dao {
	daoConfig.Do(func() {
		log.NewEntry("conf").Debug("init DaoConfig...")
		err := Config().Sub("dao").Unmarshal(&daoConfig.dao)
		if err != nil {
			log.NewEntry("conf").Panic("init DaoConfig error", log.Field{
				Key:   "err",
				Value: err,
			})
		}
		log.NewEntry("conf").Debug("init DaoConfig success")
	})
	return daoConfig.dao
}
