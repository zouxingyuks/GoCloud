package user

import (
	"GoCloud/pkg/dao"
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

	return serializer.Response{
		Code: 200,
		//todo 思考此处的msg是否需要修改
		Msg: "激活成功",
	}
}
