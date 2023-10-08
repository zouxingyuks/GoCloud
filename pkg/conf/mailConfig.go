package conf

import (
	"log"
	"sync"
)

type mail struct {
	Type string
	once sync.Once
	//同时最大发送数量
	Size int
	Smtp smtp
}

type smtp struct {
	Host string
	once sync.Once
	//端口
	Port int
	//用户名
	User string
	//密码
	Password string
	//发送者名
	Name string
	//发送者地址
	Address string
	//回复地址
	ReplyTo string
	//是否启用加密
	Encryption bool
	//SMTP 连接保留时长
	Keepalive int
}

var (
	mailConfig = new(mail)
	smtpConfig = new(smtp)
)

// MailConfig 邮件配置
func MailConfig() *mail {
	mailConfig.once.Do(func() {
		log.Printf("init mailConfig...")
		Config().Sub("mail").Unmarshal(&mailConfig)
		log.Printf("init mailConfig...end")
	})
	return mailConfig
}

// SMTPConfig SMTP配置
func SMTPConfig() *smtp {
	smtpConfig.once.Do(func() {
		log.Printf("init smtpConfig...")
		smtpConfig = &MailConfig().Smtp
		log.Printf("init smtpConfig...end")
	})
	return smtpConfig
}
