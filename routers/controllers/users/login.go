package users

import (
	"GoCloud/dao"
	"GoCloud/pkg/crypto"
	"GoCloud/pkg/log"
	"GoCloud/pkg/util"
	"GoCloud/service/serializer"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginParam 登录参数
type LoginParam struct {
	Email    string `form:"userName" json:"email" binding:"required,email" example:"test@emali.com"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20" example:"12345678"`
}

const (
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
// @Failure 404 {object} serializer.Response “用户不存在”
// @Failure 423 {object} serializer.Response "用户等待激活|用户被封禁"
// @Failure 429 {object} serializer.Response "请求过于频繁"
// @Failure 500 {object} serializer.Response "服务器错误"
// @Router /users/session [post]
func Login(c *gin.Context) {
	entry := log.NewEntry("controller.user.login")
	param := LoginParam{}
	// 1. 参数校验
	err := c.ShouldBindJSON(&param)
	if err != nil {
		res := serializer.NewResponse(entry, http.StatusBadRequest, serializer.WithMsg(ParamErrMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	//2. 用户信息获取与存在性校验
	//先在缓存中用email查找uuid,如果没有再去数据库中查找，这主要是为了防止在反复登陆时频繁访问数据库（虽然进行了接口的流量限制）
	u, err := dao.GetUserByEmail(param.Email)
	if err != nil {
		res := serializer.NewResponse(entry, http.StatusNotFound, serializer.WithMsg(CheckErrMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	//3. 不允许重复登陆
	uuid := util.Session().Get(c, "uuid")
	if uuid != nil && uuid.(string) == u.UUID {
		res := serializer.NewResponse(entry, LoginAlready, serializer.WithMsg(LoginAlreadyMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	//4. 密码校验
	if authOK, err := crypto.NewCrypto(crypto.PasswordCrypto).Check([]byte(param.Password), []byte(u.Password)); !authOK {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg(CheckErrMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	//4. 用户状态校验
	if result, response := StatusCheck(*u); !result {
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
	//准备清洗后的用户信息
	user, err := serializer.BuildUser(*u)
	if err != nil {
		res := serializer.NewResponse(entry, serializer.ServerErr, serializer.WithMsg(serializer.ServerErrMsg), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	//7. 登陆成功，清空并设置session
	util.Session().Set(c, map[string]interface{}{
		"uuid": u.UUID,
	})

	res := serializer.NewResponse(entry, http.StatusOK, serializer.WithData(user))
	c.JSON(res.Code, res)
	return
}
