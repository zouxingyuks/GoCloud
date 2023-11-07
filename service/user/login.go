package user

import (
	"GoCloud/pkg/crypto"
	"GoCloud/pkg/dao"
	"GoCloud/pkg/log"
	"GoCloud/pkg/util"
	"GoCloud/service/serializer"
	"github.com/gin-gonic/gin"
)

// 可能出现安全漏洞的地方就是sql 的部分

func (p *Param) Login(c *gin.Context) serializer.Response {
	entry := log.NewEntry("service.user")
	//账号校验
	u, err := dao.GetUser(dao.WithEmail(p.Email))
	if err != nil {
		return serializer.NewResponse(entry, 404, serializer.WithMsg("账号不存在"), serializer.WithErr(err))
	}
	//密码校验
	if authOK, err := crypto.NewCrypto(crypto.PasswordCrypto).Check([]byte(p.Password), []byte(u.Password)); !authOK {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("身份验证失败"), serializer.WithErr(err))
	}
	//激活状态校验
	if result, response := StatusCheck(u); !result {
		return response
	}

	//todo 二步验证
	//if expectedUser.TwoFactor != "" {
	//	// 需要二步验证
	//	util.SetSession(c, map[string]interface{}{
	//		"2fa_user_id": expectedUser.ID,
	//	})
	//	return serializer.Response{Code: 203}
	//}

	//todo 以及登录了就不要重复登录
	util.Session().Get(c, "user_id")
	//登陆成功，清空并设置session
	util.Session().Set(c, map[string]interface{}{
		"user_id": u.UUID,
	})
	//在此处返回基本用户数据
	user, err := serializer.BuildUser(u)
	if err != nil {
		return serializer.NewResponse(entry, 500, serializer.WithMsg("服务器错误"), serializer.WithErr(err))
	}
	return serializer.NewResponse(entry, 200, serializer.WithData(user))

}
