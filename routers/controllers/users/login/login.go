package login

import (
	"GoCloud/dao"
	"GoCloud/pkg/crypto"
	"GoCloud/pkg/log"
	serializer2 "GoCloud/pkg/serializer"
	"GoCloud/pkg/util"
	"GoCloud/routers/controllers/users/active/status"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Param 登录参数
type Param struct {
	Email    string `form:"userName" json:"email" binding:"required,email" example:"test@emali.com"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=20" example:"12345678"`
}

const (
	Already     = 201
	Msg         = "请勿重复登陆"
	ParamErrMsg = "参数错误"
	CheckErrMsg = "身份验证失败"
)

// Login 用于用户登录的接口。
// @Summary 用户登录接口
// @Description 用户登录账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param user body Param true "用户登录信息"
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
	param := Param{}
	// 1. 参数校验
	err := c.ShouldBindJSON(&param)
	if err != nil {
		entry.Warn("参数绑定错误", log.Field{
			Key:   "err",
			Value: err,
		})
		c.String(http.StatusBadRequest, "参数绑定错误")
		return
	}
	if checkParam(c, param) == false {
		entry.Warn(param.Email+"尝试登陆，但是参数错误", log.Field{
			Key:   "ip",
			Value: c.ClientIP(),
		})
		//此处不记录密码，主要是为了防止日志泄露，以及逻辑上不应该记录密码
		c.JSON(http.StatusBadRequest, "参数非法")
		return
	}

	//2. 用户信息获取与存在性校验
	//先在缓存中用email查找uuid,如果没有再去数据库中查找，这主要是为了防止在反复登陆时频繁访问数据库（虽然进行了接口的流量限制）
	u, err := dao.GetUserByEmail(param.Email)
	if err != nil {
		res := serializer2.NewResponse(entry, http.StatusNotFound, serializer2.WithMsg(CheckErrMsg), serializer2.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	//3. 不允许重复登陆
	uuid := util.Session().Get(c, "uuid")
	if uuid != nil && uuid.(string) == u.UUID {
		res := serializer2.NewResponse(entry, Already, serializer2.WithMsg(Msg), serializer2.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	//4. 密码校验
	if authOK, err := crypto.NewCrypto(crypto.PasswordCrypto).Check([]byte(param.Password), []byte(u.Password)); !authOK {
		res := serializer2.NewResponse(entry, 400, serializer2.WithMsg(CheckErrMsg), serializer2.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	//4. 用户状态校验
	if result, response := status.Check(*u); !result {
		c.JSON(response.Code, response)
		return
	}

	//准备清洗后的用户信息
	user, err := serializer2.BuildUser(*u)
	if err != nil {
		res := serializer2.NewResponse(entry, serializer2.ServerErr, serializer2.WithMsg(serializer2.ServerErrMsg), serializer2.WithErr(err))
		c.JSON(res.Code, res)
		return
	}

	//7. 登陆成功，清空并设置session
	util.Session().Set(c, map[string]interface{}{
		"uuid": u.UUID,
	})

	res := serializer2.NewResponse(entry, http.StatusOK, serializer2.WithData(user))
	c.JSON(res.Code, res)
	return
}
