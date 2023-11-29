package conf

import (
	"GoCloud/pkg/log"
	cachePkg "GoCloud/service/cache"
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
}
type cache struct {
	Kind  cachePkg.Kind
	Redis redis
}
type dao struct {
	Database database
	Cache    cache
}

var daoConfig = new(struct {
	*dao
	sync.Once
})

func DaoConfig() *dao {
	daoConfig.Do(func() {
		log.NewEntry("conf").Info("init DaoConfig...")
		err := Config().Sub("dao").Unmarshal(&daoConfig.dao)
		if err != nil {
			log.NewEntry("conf").Panic("init DaoConfig error", log.Field{
				Key:   "err",
				Value: err,
			})
		}
		log.NewEntry("conf").Info("init DaoConfig success")
	})
	return daoConfig.dao
}
