package register

import (
	"GoCloud/conf"
	"GoCloud/dao"
	"GoCloud/pkg/crypto"
	"GoCloud/pkg/log"
	"GoCloud/pkg/rbac"
	"GoCloud/pkg/serializer"
	"GoCloud/routers/controllers/users/active"
	"GoCloud/routers/controllers/users/active/status"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 设计思想：
// 1. 防止箭头形代码：提高代码可读性

type Param struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

// Register 用于用户注册的接口。
// @Summary 用户注册接口
// @Description 用户注册新账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param user body Param true "用户注册信息"
// @Success 200 {object} serializer.Response "注册成功"
// @Failure 400 {object} serializer.Response "参数错误"
// @Failure 500 {object} serializer.Response "服务异常"
// @Router /users [post]
func Register(c *gin.Context) {
	entry := log.NewEntry("controller.user.register")

	// 1. 检查注册功能是否开启
	if !conf.ServiceConfig().User.RegisterEnable {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg("注册功能已关闭"))
		c.JSON(res.Code, res)
		return
	}

	param := Param{}
	// 2.参数绑定
	err := c.ShouldBindJSON(&param)
	if err != nil {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg("参数错误"), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	// 2. 检查
	if !checkParam(c, param) {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg("参数错误"))
		c.JSON(res.Code, res)
		return
	}

	// 3. 生成用户信息
	user := dao.NewUser(
		dao.WithEmail(param.Email),
		// 默认用户状态为未激活
		dao.WithStatus(dao.UserNotActivated),
	)
	// 进行密码加密
	user.Password, err = crypto.NewCrypto(crypto.PasswordCrypto).Encrypt([]byte(param.Password))
	if err != nil {
		res := serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(errors.Wrap(err, "密码加密失败")))
		c.JSON(res.Code, res)
		return
	}

	//  如果不需要邮箱验证，则设置账户为激活模式
	if !conf.ServiceConfig().User.EmailVerify {
		user.Status = dao.UserActive
	}

	// 4. 检测邮箱是否已经注册
	var tUser *dao.User
	tUser, err = dao.GetUserByEmail(user.Email)
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		// 如果已经注册，检查用户状态
		_, res := status.Check(*tUser)
		c.JSON(res.Code, res)
		return
	}

	// 5. 创建用户
	err = dao.CreateUser(user)
	if err != nil {
		res := serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	// 分配默认权限
	_, err = rbac.ChangeRoleForUser(user.UUID, rbac.User)
	if err != nil {
		res := serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	// 6. 发送激活邮件
	if user.Status == dao.UserNotActivated {
		// 发送激活邮件
		err = active.SendActivationEmail(user.UUID, user.Email)
		if err != nil {
			res := serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
			c.JSON(res.Code, res)
			return
		}
		res := serializer.NewResponse(entry, 200, serializer.WithMsg("注册成功,请前往邮箱激活"))
		c.JSON(res.Code, res)
		return
	}
	res := serializer.NewResponse(entry, 200, serializer.WithMsg("注册成功"))
	c.JSON(res.Code, res)
	return
}
