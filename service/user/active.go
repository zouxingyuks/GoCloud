package user

import (
	"GoCloud/pkg/dao"
	log2 "GoCloud/pkg/log"
	"GoCloud/pkg/serializer"
)

// Active  管理用户注册服务
func (p *Param) Active(token string) serializer.Response {

	//todo 1. 校验token，寻找对应的user
	user := dao.User{}

	//todo 2. 激活用户
	if user.Status != dao.UserNotActivated {
		return serializer.Response{
			//todo 思考此处的code是否需要修改
			Code: 400,
			Msg:  "用户已激活",
		}
	}
	user.Status = dao.UserActive
	//todo 3. 返回结果

	return serializer.NewResponse(log2.NewEntry("service.user"), 200, "用户激活成功")
}
func StatusCheck(u dao.User) (bool, serializer.Response) {
	entry := log2.NewEntry("service.user")

	switch u.Status {
	//未激活状态
	case dao.UserNotActivated:
		//未激活是用户的问题
		//todo 重新发送激活邮件
		return false, serializer.NewResponse(entry, 400, "账户等待激活,已经重新发送激活邮件", nil)
	//激活状态
	case dao.UserActive:
		return true, serializer.Response{}
	//未知状态
	default:
		return false, serializer.NewResponse(entry, 400, "账户状态异常", nil)
	}

}
