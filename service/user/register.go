package user

import (
	"GoCloud/models"
	"GoCloud/pkg/conf"
	"GoCloud/pkg/serializer"
	"GoCloud/pkg/util/filter"
	"crypto/sha256"
	"fmt"
	"net/http"
)

// todo 是否有必要记录 注册用户记录
func (p *Param) Register() serializer.Response {
	if !filter.Facade.IsValidEmail(p.Email) {
		return serializer.Response{
			Code: 400,
			Msg:  serializer.EmailInvalided.Error(),
		}
	}
	//todo 邮箱格式校验

	//todo 加载更多模块
	//// 相关设定
	//此类变量应该是函数开始时就获取，避免临时修改导致的错误
	//options := model.GetSettingByNames("email_active")
	//isEmailRequired := models.IsTrueVal(options["email_active"])
	defaultGroup := conf.SystemConfig.Get(conf.UserConfigDefaultGroup)

	// 创建新的用户对象
	user := models.NewUser()
	user.Email = p.Email
	// 默认用户名为邮箱前缀
	user.UserName = defaultUserName(p.Email)
	err := user.SetPassword(p.Password)
	if err != nil {
		return serializer.Response{
			Code: 400,
			Msg:  serializer.PassWordInvalided.Error(),
		}
	}
	user.Status = models.UserActive
	//if isEmailRequired {
	//	user.Status = model.UserNotActivated
	//}
	user.GroupID = uint(defaultGroup.(int))
	//todo 取消测试
	//userNotActivated := false
	//todo 补全数据库创建逻辑
	//先进行尝试创建
	if err := models.DB().Create(&user).Error; err != nil {
		//创建失败后进一步判断错误情况

		//检查已存在使用者是否尚未激活
		expectedUser, err := models.GetUserByEmail(user.Email)
		//如若尚未激活，则将用户状态设置为未激活
		if expectedUser.Status == models.UserNotActivated {
			//todo 未激活用户的处理
			//userNotActivated = true
			user = expectedUser
		} else {
			//
			return serializer.Err(http.StatusBadRequest, serializer.EmailExist.Error(), err)
		}
	}
	// todo 激活邮件
	// 发送激活邮件
	//if isEmailRequired {
	//
	//	// 签名激活请求API
	//	base := model.GetSiteURL()
	//	userID := hashid.HashID(user.ID, hashid.UserID)
	//	controller, _ := url.Parse("/api/v3/user/activate/" + userID)
	//	activateURL, err := auth.SignURI(auth.General, base.ResolveReference(controller).String(), 86400)
	//	if err != nil {
	//		return serializer.Err(serializer.CodeEncryptError, "Failed to sign the activation link", err)
	//	}
	//
	//	// 取得签名
	//	credential := activateURL.Query().Get("sign")
	//
	//	// 生成对用户访问的激活地址
	//	controller, _ = url.Parse("/activate")
	//	finalURL := base.ResolveReference(controller)
	//	queries := finalURL.Query()
	//	queries.Add("id", userID)
	//	queries.Add("sign", credential)
	//	finalURL.RawQuery = queries.Encode()
	//
	//	// 返送激活邮件
	//	title, body := email.NewActivationEmail(user.Email,
	//		finalURL.String(),
	//	)
	//	if err := email.Send(user.Email, title, body); err != nil {
	//		return serializer.Err(serializer.CodeFailedSendEmail, "Failed to send activation email", err)
	//	}
	//	if userNotActivated == true {
	//		//原本在上面要抛出的DBErr，放来这边抛出
	//		return serializer.Err(serializer.CodeEmailSent, "User is not activated, activation email has been resent", nil)
	//	} else {
	//		return serializer.Response{Code: 203}
	//	}
	//todo 若是设置了激活邮件，则需要进行激活
	//return serializer.Response{
	//	Code:  200,
	//	Msg:   WaitActive,
	//}
	//}
	//todo 继续编写注册器
	return serializer.Response{
		Code: 200,
		Msg:  serializer.RegisterSucceed.Error(),
	}
}
func defaultUserName(email string) string {
	// 使用SHA-256哈希算法生成摘要
	hash := sha256.Sum256([]byte(email))
	// 将摘要转换为字符串表示
	result := fmt.Sprintf("%x", hash)
	return result
}
