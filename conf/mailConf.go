package conf

import (
	"GoCloud/pkg/email/model"
	"GoCloud/pkg/log"
)

var (
	mailConfig = new(model.Mail)
	smtpConfig = new(model.SMTP)
)

// MailConfig 邮件配置
func MailConfig() *model.Mail {
	mailConfig.Once.Do(func() {
		log.NewEntry("conf").Info("init mailConfig...")
		Config().Sub("mail").Unmarshal(&mailConfig)
		log.NewEntry("conf").Info("init mailConfig...end")
	})
	return mailConfig
}

// SMTPConfig SMTP配置
func SMTPConfig() *model.SMTP {
	smtpConfig.Once.Do(func() {
		log.NewEntry("conf").Info("init smtpConfig...")
		smtpConfig = &MailConfig().Smtp
		log.NewEntry("conf").Info("init smtpConfig...end")
	})
	return smtpConfig
}
