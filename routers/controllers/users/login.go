package users

import (
	"GoCloud/dao"
	"GoCloud/pkg/crypto"
	"GoCloud/pkg/log"
	"GoCloud/pkg/util"
	"GoCloud/service/serializer"
	"github.com/gin-gonic/gin"
)

// LoginParam 登录参数
type LoginParam struct {
	Email    string `json:"email" binding:"required,email" example:"test@emali.com""`
	Password string `json:"password" binding:"required,min=8,max=20" example:"12345678"`
}

const (
	LoginSuccess    = 200
	LoginAlready    = 201
	LoginAlreadyMsg = "请勿重复登陆"
	ParamErrMsg     = "参数错误"
	CheckErrMsg     = "身份验证失败"
)

// Login 用于用户登录的接口。
// @Summary 用户登录接口
// @Description 用户登录账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param user body LoginParam true "用户登录信息"
// @Success 200 {object} serializer.Response "登录成功"
// @Success 201 {object} serializer.Response "请勿重复登陆"
// @Failure 400 {object} serializer.Response "参数错误"
// @Failure 404 {object} serializer.Response “身份验证失败”
// @Router /users/session [post]
func Login(c *gin.Context) {
	entry := log.NewEntry("controller.user.login")
	param := LoginParam{}
	// 1. 参数校验
	err := c.ShouldBindJSON(&param)
	if err != nil {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg(ParamErrMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	//2. 用户信息获取与存在性校验
	//先在缓存中用email查找uuid,如果没有再去数据库中查找，这主要是为了防止在反复登陆时频繁访问数据库（虽然进行了接口的流量限制）
	uuid := ""
	iUUID, err := dao.Cache().Get("email:" + param.Email)
	if err == nil {
		uuid = iUUID.(string)
	}
	//账号校验
	u, err := dao.GetUser(dao.WithEmail(param.Email), dao.WithUUID(uuid))
	if err != nil {
		// 此处之所以不指出是账号错误还是密码错误，是为了防止账号被暴力破解
		res := serializer.NewResponse(entry, 400, serializer.WithMsg(ParamErrMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	//3. 密码校验
	if authOK, err := crypto.NewCrypto(crypto.PasswordCrypto).Check([]byte(param.Password), []byte(u.Password)); !authOK {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg(CheckErrMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	//4. 用户状态校验
	if result, response := StatusCheck(u); !result {
		c.JSON(response.Code, response)
		return
	}

	//5. 二步验证
	//if expectedUser.TwoFactor != "" {
	//	// 需要二步验证
	//	util.SetSession(c, map[string]interface{}{
	//		"2fa_user_id": expectedUser.ID,
	//	})
	//	return serializer.Response{Code: 203}
	//}

	//6. 不允许重复登陆
	iUUID = util.Session().Get(c, "uuid")
	if iUUID != nil && iUUID.(string) == u.UUID {
		res := serializer.NewResponse(entry, LoginAlready, serializer.WithMsg(LoginAlreadyMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
	}

	//7. 登陆成功，清空并设置session
	util.Session().Set(c, map[string]interface{}{
		"uuid": u.UUID,
	})
	//返回清洗后的用户信息
	user, err := serializer.BuildUser(u)
	if err != nil {
		res := serializer.NewResponse(entry, serializer.ServerErr, serializer.WithMsg(serializer.ServerErrMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	res := serializer.NewResponse(entry, LoginSuccess, serializer.WithData(user))
	c.JSON(res.Code, res)
	return
}
