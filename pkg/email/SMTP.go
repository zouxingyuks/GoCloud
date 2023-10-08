package email

import (
	"GoCloud/pkg/conf"
	"GoCloud/pkg/log"
	"context"
	"fmt"
	"github.com/go-mail/mail"
	"github.com/panjf2000/ants/v2"
	"github.com/pkg/errors"
	"sync"
)

type SMTP struct {
	pool *ants.Pool
	log.IEntry
	context context.Context
	wg      sync.WaitGroup
}

func NewSMTP() *SMTP {
	var (
		err  error
		smtp = new(SMTP)
	)
	//初始化日志
	smtp.IEntry = log.NewEntry("email.smtp")
	//初始化一个协程池
	smtp.pool, err = ants.NewPool(conf.MailConfig().Size)
	if err != nil {
		smtp.Panic(errors.Wrap(err, "初始化协程池失败").Error())
		return nil
	}
	//todo 监控配置文件变化
	//smtp.watch()
	return smtp
}

// todo 检查是否需要加锁
func (S *SMTP) watch() {
	go func() {
		for {
			select {
			case <-S.context.Done():
				S.Info("SMTP协程池已关闭")
				return
			case <-conf.Change:
				S.pool.Tune(conf.MailConfig().Size)
				S.Info(fmt.Sprintf("SMTP协程池大小已调整为%d", conf.MailConfig().Size))
			}

		}
	}()
}

// Send 发送邮件功能
func (S *SMTP) send(to, title, body string) error {
	m := mail.NewMessage()
	m.SetAddressHeader("From", conf.MailConfig().Smtp.Address, conf.MailConfig().Smtp.Name)
	m.SetAddressHeader("Reply-To", conf.MailConfig().Smtp.ReplyTo, conf.MailConfig().Smtp.Name)
	m.SetHeader("To", to)
	m.SetHeader("Subject", title)
	m.SetBody("text/html", body)
	d := mail.NewDialer(conf.MailConfig().Smtp.Host, conf.MailConfig().Smtp.Port, conf.MailConfig().Smtp.User, conf.MailConfig().Smtp.Password)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return errors.Wrap(err, "发送邮件失败")
	}
	//todo 编辑记录信息
	S.Info(fmt.Sprintf("发送邮件给%s成功", to))
	return nil
}

func (S *SMTP) Submit(to, title, body string) error {
	S.wg.Add(1)
	task := func() {
		err := S.send(to, title, body)
		//此问题不影响发送邮件，但是值得关注
		if err != nil {
			S.Warn(err.Error())
		}
		S.wg.Done()
	}
	err := S.pool.Submit(task)
	if err != nil {
		//因为提交失败了，所以不需要等待
		S.wg.Done()
		return errors.Wrapf(err, "提交任务失败")
	}
	return nil

}
func (S *SMTP) Close() {
	S.wg.Wait()
	S.pool.Release()
}
