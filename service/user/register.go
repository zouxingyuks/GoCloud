package user

import (
	"GoCloud/conf"
	"GoCloud/dao"
	"GoCloud/pkg/log"
	"GoCloud/pkg/util/filter"
	"GoCloud/service/rbac"
	"GoCloud/service/serializer"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// sql 防注入： 使用 orm 进行数据库操作，orm 会自动对 sql 进行预编译，防止 sql 注入。

// Register  管理用户注册服务
func (p *Param) Register() serializer.Response {
	entry := log.NewEntry("service.user.register")
	if !conf.ServiceConfig().User.RegisterEnable {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("注册功能已关闭"))
	}
	if !filter.Facade.IsValidEmail(p.Email) {
		//行为日志
		return serializer.NewResponse(entry, 400, serializer.WithMsg("邮箱非法"), serializer.WithField(
			log.Field{
				Key:   "Email",
				Value: p.Email,
			}))
	}

	// 相关设定
	//此类变量应该是函数开始时就获取，避免竟态条件导致的错误
	isEmailRequired := conf.ServiceConfig().User.EmailVerify

	// 创建新的用户对象
	user := dao.NewUser(
		dao.WithEmail(p.Email),
		// 默认用户状态为未激活
		dao.WithStatus(dao.UserNotActivated),
	)
	// 设置密码
	err := user.SetPassword(p.Password)
	if err != nil {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("密码非法"))
	}
	//  如果不需要邮箱验证，则设置账户为激活模式
	if !isEmailRequired {
		user.Status = dao.UserActive
	}
	var tUser dao.User
	tUser, err = dao.GetUser(dao.WithEmail(user.Email))
	// 检测邮箱是否已经注册
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err := dao.CreateUser(&user)
		if err != nil {
			return serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
		}
	} else {
		// 检查用户状态
		if result, response := StatusCheck(tUser); !result {
			return response
		} else {
			return serializer.NewResponse(entry, 400, serializer.WithMsg("Email already in use"), serializer.WithErr(err))

		}
	}
	//分配角色,默认为user
	res := rbac.AssignRolesToUser(user.UUID, "user")
	if res.Code != 200 {
		return serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(errors.New(res.Msg)))
	}
	return serializer.NewResponse(entry, 200, serializer.WithMsg("注册成功"))
}
