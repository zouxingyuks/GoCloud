package routers

import (
	"GoCloud/pkg/conf"
	"GoCloud/pkg/log"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	log.NewEntry("router").Info("当前运行模式为" + conf.SystemConfig().Mode)
	if conf.SystemConfig().Mode == "master" {
		return InitMasterRouter()
	}
	return InitSlaveRouter()
}

// InitSlaveRouter 初始化从机模式路由
func InitSlaveRouter() *gin.Engine {
	r := gin.Default()
	//todo 从机模式
	return r
}

// InitMasterRouter 初始化主机模式路由
func InitMasterRouter() *gin.Engine {
	router := newApi(0)
	r := router.load()
	return r
}
func InitCORS(router *gin.Engine) {
	//if conf.CORSConfig.AllowOrigins[0] != "UNSET" {
	//	router.Use(cors.New(cors.config{
	//		AllowOrigins:     conf.CORSConfig.AllowOrigins,
	//		AllowMethods:     conf.CORSConfig.AllowMethods,
	//		AllowHeaders:     conf.CORSConfig.AllowHeaders,
	//		AllowCredentials: conf.CORSConfig.AllowCredentials,
	//		ExposeHeaders:    conf.CORSConfig.ExposeHeaders,
	//	}))
	//	return
	//}
	//
	//// slave模式下未启动跨域的警告
	//if conf.SystemConfig.Mode == "slave" {
	//	util.Log().Warning("You are running Cloudreve as slave node, if you are using slave storage policy, please enable CORS feature in config file, otherwise file cannot be uploaded from Master site.")
	//}
}
