package user

import (
	"GoCloud/pkg/conf"
	"GoCloud/pkg/dao"
	log2 "GoCloud/pkg/log"
	"GoCloud/pkg/serializer"
	"GoCloud/pkg/util/filter"
	"crypto/sha256"
	"fmt"
)

// Register  管理用户注册服务
func (p *Param) Register() serializer.Response {
	entry := log2.NewEntry("service.user.register")
	if !filter.Facade.IsValidEmail(p.Email) {
		//行为日志
		return serializer.NewResponse(entry, 400, serializer.WithMsg("邮箱非法"), serializer.WithField(
			log2.Field{
				Key:   "Email",
				Value: p.Email,
			}))
	}

	// 相关设定
	//此类变量应该是函数开始时就获取，避免临时修改导致的错误
	isEmailRequired := conf.UserConfig().EmailVerify
	// 创建新的用户对象
	user := dao.NewUser()
	user.Email = p.Email
	// 默认用户名为邮箱前缀
	user.NickName = defaultUserName(p.Email)
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
	//先进行尝试创建
	if err := dao.DB().Create(&user).Error; err != nil {
		//创建失败后进一步判断错误情况

		//检查此账户状态
		expectedUser, err := dao.GetUser(dao.WithEmail(user.Email))
		//如若尚未激活，则将用户状态设置为未激活
		if result, response := StatusCheck(expectedUser); err == nil && !result {
			return response
		} else if err != nil {
			return serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
		} else {
			return serializer.NewResponse(entry, 400, serializer.WithMsg("Email already in use"), serializer.WithErr(err))
		}
	}
	return serializer.NewResponse(entry, 200, serializer.WithMsg("注册成功"))
}
func defaultUserName(email string) string {
	// 使用SHA-256哈希算法生成摘要
	hash := sha256.Sum256([]byte(email))
	// 将摘要转换为字符串表示
	result := fmt.Sprintf("小可爱%x", hash)
	return result
}
