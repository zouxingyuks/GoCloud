package conf

import "GoCloud/pkg/log"

func ReloadConfig() {
	//todo 考虑是不是全部改成把  conf 作为最底层依赖包
	//在此处更新一些配置
	{
		//有些配置是需要手动更改的，比如并发数之类的
		temp := Config().GetBool("usercontroller.emailverify")
		if temp {
			log.NewEntry("conf").Info("userConfig.EmailVerify Changed", log.Field{
				Key: "userConfig.EmailVerify",
			})
			serviceConfig.User.EmailVerify = temp
		}
	}
	// debug 可能导致日志输出位置发生变化，因此最后更改
	{
		temp := Config().GetBool("system.debug")
		if temp != systemConfig.Debug {
			log.NewEntry("conf").Info("systemConfig.Debug Changed", log.Field{
				Key:   "systemConfig.Debug",
				Value: temp,
			})
			systemConfig.Debug = temp
			log.SetDebug(systemConfig.Debug)
		}
	}
	//todo 更新一个广播功能
	//Change <- struct{}{}
}
