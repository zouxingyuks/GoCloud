package user

import (
	"GoCloud/pkg/conf"
	"GoCloud/pkg/dao"
	"GoCloud/pkg/log"
	"GoCloud/pkg/util/filter"
	"GoCloud/service/rbac"
	"GoCloud/service/serializer"
	"crypto/sha256"
	"fmt"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// Register  管理用户注册服务
func (p *Param) Register() serializer.Response {
	entry := log.NewEntry("service.user.register")
	if !filter.Facade.IsValidEmail(p.Email) {
		//行为日志
		return serializer.NewResponse(entry, 400, serializer.WithMsg("邮箱非法"), serializer.WithField(
			log.Field{
				Key:   "Email",
				Value: p.Email,
			}))
	}

	// 相关设定
	//todo 避免竟态漏洞，考虑是否存在，如何防范
	//此类变量应该是函数开始时就获取，避免竟态条件导致的错误
	isEmailRequired := conf.UserConfig().EmailVerify

	// 创建新的用户对象
	user := dao.NewUser(
		dao.WithEmail(p.Email),
		// 默认用户名为邮箱前缀
		dao.WithNickName(defaultUserName(p.Email)))
	// 设置密码
	err := user.SetPassword(p.Password)
	if err != nil {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("密码非法"))
	}
	//  如果不需要邮箱验证，则设置账户为激活模式
	if !isEmailRequired {
		user.Status = dao.UserActive
	}
	//todo 继续测试此部分
	//todo 对应 sql 的防注入

	var tUser dao.User
	tUser, err = dao.GetUser(dao.WithEmail(user.Email))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		//cErr := dao.DB().Create(&user).Error
		err := dao.CreateUser(&user)
		log.NewEntry("service.user.register").Debug("create user", log.Field{
			Key:   "user",
			Value: user,
		}, log.Field{
			Key:   "err",
			Value: err,
		})
		if err != nil {
			return serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
		}
	} else {
		//如若尚未激活，则将用户状态设置为未激活
		if result, response := StatusCheck(tUser); err == nil && !result {
			return response

		} else {
			return serializer.NewResponse(entry, 400, serializer.WithMsg("Email already in use"), serializer.WithErr(err))

		}
		//todo 对于其他错误的处理
		//else {
		//	return serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err), serializer.WithData(errors.Is(err, gorm.ErrDuplicatedKey)))
		//
		//}
	}
	//分配角色,默认为user
	rbac.AssignRolesToUser(user.UUID, "user")
	return serializer.NewResponse(entry, 200, serializer.WithMsg("注册成功"))
}
func defaultUserName(email string) string {
	// 使用SHA-256哈希算法生成摘要
	hash := sha256.Sum256([]byte(email))
	// 将摘要转换为字符串表示
	result := fmt.Sprintf("小可爱%x", hash)
	return result
}
