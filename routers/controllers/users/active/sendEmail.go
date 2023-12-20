package active

import (
	"GoCloud/pkg/log"
	"GoCloud/service/email"
	"bytes"
	"github.com/pkg/errors"
	"html/template"
)

// SendActivationEmail 创建一个新的激活邮件
func SendActivationEmail(uuid, To string) error {
	var (
		tpl  bytes.Buffer
		err  error
		data = ActivationEmailData{
			To: To,
		}
	)
	// 生成激活链接
	data.ActivationLink, err = generateURL(uuid)
	if err != nil {
		return err
	}

	// 创建一个新的模板实例，并解析模板内容
	emailTml := defaultActivationEmailTmpl
	tmpl, err := template.New("email").Parse(emailTml)
	if err != nil {
		return errors.Wrap(err, "Error parsing email template")
	}
	if err := tmpl.Execute(&tpl, data); err != nil {
		return errors.Wrapf(err, "Error rendering email template with data %+v", data)
	}
	// 模板渲染结果
	result := tpl.String()
	err = email.Driver().Submit(To, ActivationEmailTitle, result)
	log.NewEntry("service.user.sendActivationEmail").Info("发送激活邮件", log.Field{
		Key:   "To",
		Value: To,
	})
	return err
}

type ActivationEmailData struct {
	// To 收件人邮箱
	To string
	// ActivationLink 激活链接
	ActivationLink string
}
