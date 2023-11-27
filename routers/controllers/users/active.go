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
	//// 2.解析token
	//token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
	//	return []byte(conf.StaticConfig().Session.Secret), nil
	//})
	//if err != nil {
	//	res := serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
	//
	//	c.JSON(res.Code, res)
	//	return
	//}
	//
	//// 3. 校验token
	//// 3.1.校验token是否有效
	//// 3.2.校验token中的uuid是否有效
	//// 3.3.校验token中的act是否为active
	//if claims, ok := token.Claims.(jwt.MapClaims); token.Valid && ok && claims["sub"] != nil && claims["act"] == "active" {
	//	uuid := claims["sub"].(string)
	//	_, err = dao.SetUser(dao.WithUUID(uuid), dao.WithStatus(dao.UserActive))
	//	if err != nil {
	//		res := serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
	//		c.JSON(res.Code, res)
	//		return
	//	}
	//	res := serializer.NewResponse(entry, 200, serializer.WithMsg("激活成功"))
	//	c.JSON(res.Code, res)
	//	return
	//} else {
	//	res := serializer.NewResponse(entry, 401, serializer.WithMsg("无效的token"))
	//	c.JSON(res.Code, res)
	//	return
	//}
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

const defaultActivationEmailTmpl = `
<!DOCTYPE html>
<html>
<head>
    <title>账户激活</title>
</head>
<body style="background-color: #F3F4F6; padding: 5rem; display: flex; justify-content: center; align-items: center;">
<div style="background-color: #ffffff; padding: 2rem; border-radius: 0.5rem; box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1); width: 100%; max-width: 400px;">
    <h1 style="font-size: 1.5rem; font-weight: 600; margin-bottom: 1rem; color: #2563EB;">欢迎加入 GoCloud！</h1>
    <p style="color: #4B5563; margin-bottom: 1rem;">感谢您注册 GoCloud。请点击下面的链接来激活您的账户：</p>
    <a href="{{.ActivationLink}}" target="_blank"
       style="display: inline-block; background-color: #2563EB; color: #ffffff; padding: 0.5rem 1rem; border-radius: 0.25rem; text-decoration: none;">点击这里激活账户</a>
    <p style="color: #4B5563; margin-top: 1rem; margin-bottom: 0.5rem;">或者复制下面的链接到浏览器地址栏：</p>
    <p style="color: #4B5563; margin-bottom: 1rem; word-break: break-all;">{{.ActivationLink}}</p>
    <p style="color: #4B5563; margin-bottom: 0.5rem;">如果您没有注册 GoCloud，请忽略此邮件。</p>
    <p style="color: #4B5563;">谢谢！</p>
    <p style="color: #4B5563; margin-top: 1rem;">GoCloud 团队</p>
</div>
</body>
</html>`
const ActivationEmailTitle = "GoCloud 账户激活"

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
