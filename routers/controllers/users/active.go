package users

import (
	"GoCloud/conf"
	"GoCloud/dao"
	"GoCloud/pkg/log"
	"GoCloud/service/email"
	"GoCloud/service/serializer"
	token2 "GoCloud/service/token"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"html/template"
	"net/url"
	"time"
)

const activeAction = "active"

var exp = time.Hour

// UserActivate
// @Summary 用户激活接口
// @Description 用户激活账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param token path string true "用户激活token"
// @Success 200 {object} serializer.Response "激活成功"
// @Failure 400 {object} serializer.Response "参数错误"
// @Failure 500 {object} serializer.Response "服务异常"
// @Router /users/activate/{token} [get]
func UserActivate(c *gin.Context) {
	entry := log.NewEntry("controller.user.active")

	tokenStr := c.Param("token")
	// 1.参数校验
	if tokenStr == "" {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg("参数错误"))
		c.JSON(res.Code, res)
		return
	}
	// 2.解析token
	sub, err := token2.Parse(tokenStr, activeAction)
	if err != nil {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg("无效的token"), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	// 3. 激活用户
	uuid := sub.(string)
	_, err = dao.SetUser(dao.WithUUID(uuid), dao.WithStatus(dao.UserActive))
	if err != nil {
		res := serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	res := serializer.NewResponse(entry, 200, serializer.WithMsg("激活成功"))
	c.JSON(res.Code, res)
	return
}

// SendActivationEmail 创建一个新的激活邮件
func SendActivationEmail(uuid, To string) error {
	// 创建一个新的模板实例，并解析模板内容
	emailTml := defaultActivationEmailTmpl
	tmpl, err := template.New("email").Parse(emailTml)
	if err != nil {
		return errors.Wrap(err, "Error parsing email template")
	}
	// 渲染模板
	var tpl bytes.Buffer
	var data = ActivationEmailData{
		To: To,
	}
	// 生成激活链接
	if token, err := token2.Generate(activeAction, uuid, exp); err != nil {
		return errors.Wrap(err, "Error generating activation token")
	} else {
		u := url.URL{}
		u.Path, err = url.JoinPath("/users/activate/" + token)
		if err != nil {
			return errors.Wrap(err, "Error generating activation link")
		}
		u.Scheme = "http"
		u.Host = conf.SiteConfig().Domain
		if conf.SiteConfig().SSL {
			u.Scheme = "https"
		}
		data.ActivationLink = u.String()
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
