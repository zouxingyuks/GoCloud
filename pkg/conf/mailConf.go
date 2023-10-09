package conf

import (
	"GoCloud/pkg/email/model"
	"log"
)

var (
	mailConfig = new(model.Mail)
	smtpConfig = new(model.SMTP)
)

// MailConfig 邮件配置
func MailConfig() *model.Mail {
	mailConfig.Once.Do(func() {
		log.Printf("init mailConfig...")
		Config().Sub("mail").Unmarshal(&mailConfig)
		log.Printf("init mailConfig...end")
	})
	return mailConfig
}

// SMTPConfig SMTP配置
func SMTPConfig() *model.SMTP {
	smtpConfig.Once.Do(func() {
		log.Printf("init smtpConfig...")
		smtpConfig = &MailConfig().Smtp
		log.Printf("init smtpConfig...end")
	})
	return smtpConfig
}
