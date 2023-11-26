package users

import (
	"GoCloud/dao"
	"GoCloud/pkg/log"
	"GoCloud/service/serializer"
	"fmt"
)

// StatusCheck 检查用户状态
// bool 用于检测用户是否具有登录资格，true表示可以登录，false表示不可以登录
// serializer.Response 用于返回给注册接口的响应
func StatusCheck(u dao.User) (bool, serializer.Response) {
	entry := log.NewEntry("service.user.active")
	switch u.Status {
	//未激活状态
	case dao.UserNotActivated:
		err := SendActivationEmail(u.UUID, u.Email)
		fmt.Println(err)
		if err != nil {
			return false, serializer.NewResponse(entry, 500, serializer.WithMsg("账户等待激活,邮件发送失败"), serializer.WithErr(err))
		}
		return false, serializer.NewResponse(entry, 400, serializer.WithMsg("账户等待激活,已经重新发送激活邮件"))
	//激活状态
	case dao.UserActive:
		return true, serializer.NewResponse(entry, 400, serializer.WithMsg("Email already in use"))
	//封禁状态
	case dao.UserBaned:
		return false, serializer.NewResponse(entry, 400, serializer.WithMsg("账户已被封禁"))

	default:
		return false, serializer.NewResponse(entry, 400, serializer.WithMsg("账户状态异常"))
	}

}
