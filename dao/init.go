package dao

import (
	"GoCloud/pkg/conf"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

var DB *gorm.DB

//todo 修改

// 连接数据库
func Init() {
	//util.Log().Info("Initializing database connection...")
	var (
		db  *gorm.DB
		err error
	)
	if gin.Mode() == gin.TestMode {
		// 测试模式下，使用内存数据库
		db, err = gorm.Open("sqlite", ":memory:")
	} else {
		switch conf.DatabaseConfig.Type {
		//case "UNSET", "sqlite":
		//	// 未指定数据库或者明确指定为 sqlite 时，使用 SQLite 数据库
		//	db, err = gorm.Open("sqlite", util.RelativePath(conf.DatabaseConfig.DBFile))
		//case "postgres":
		//	db, err = gorm.Open(confDBType, fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		//		conf.DatabaseConfig.Host,
		//		conf.DatabaseConfig.User,
		//		conf.DatabaseConfig.Password,
		//		conf.DatabaseConfig.Name,
		//		conf.DatabaseConfig.Port))
		case "mysql":
			db, err = connMysql()
			//default:
			//	util.Log().Panic("Unsupported database type %q.", confDBType)
			//}
		}
	}
	//todo  替换log
	panic(err)
	//util.Log().Panic("Failed to connect to database: %s", err)
	//todo 这部分没看懂
	// 处理表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return conf.DatabaseConfig.TablePrefix + defaultTableName
	}
	// Debug模式下，输出所有 SQL 日志
	if conf.SystemConfig.Debug {
		db.LogMode(true)
	} else {
		db.LogMode(false)
	}
	// 设置连接池
	//设置连接池
	db.DB().SetMaxIdleConns(50)
	if conf.DatabaseConfig.Type == "sqlite" || conf.DatabaseConfig.Type == "UNSET" {
		db.DB().SetMaxOpenConns(1)
	} else {
		db.DB().SetMaxOpenConns(100)
	}

	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db
	//执行迁移

	//todo 数据未迁移
	//migration()
}

// 连接MySQL数据库
// bash: go get gorm.io/driver/mysql
func connMysql() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", conf.DatabaseConfig.User, conf.DatabaseConfig.Password, conf.DatabaseConfig.Host, conf.DatabaseConfig.Name, conf.DatabaseConfig.Charset)
	db, err := gorm.Open("mysql", dsn)
	return db, err
}

// 连接SQLite数据库
// 需要安装 go get gorm.io/driver/sqlite
//func connSQLite() {
//	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
//	if err != nil {
//		log.Errorln("数据库连接失败")
//		panic(err)
//	}
//	DB = db
//}
