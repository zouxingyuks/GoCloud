package controllers

//进行多一层封装，以保证 api 的路径的修改方便，以及可以不用修改下层实现逻辑
import "github.com/gin-gonic/gin"

type v0 struct {
}

var V0 v0

// UserRegister 用户注册接口
// @Summary 用户注册接口
// @Description 用户注册新账户
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} serializer.Response "注册成功"
// @Failure 400 {object} serializer.Response "参数错误"
// @Router /users/register [post]
func (V0) UserRegister(c *gin.Context) {
	userRegister(c)
}

// UserLogin 用于用户登录的接口。
// @Summary 用户登录接口
// @Description 用户登录账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param user body user.Param true "用户登录信息"
// @Success 200 {object} serializer.Response "登录成功" Example({"message": "登录成功"})
// @Failure 400 {object} serializer.Response "参数错误" Example({"message": "参数错误"})
// @Router /users/session [post]
func (v0) UserLogin(c *gin.Context) {
	userLogin(c)
}

// UserActivate
// @Summary 用户激活接口
// @Description 用户激活账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param token path string true "用户激活token"
// @Success 200 {object} serializer.Response "激活成功" Example({"message": "激活成功"})
// @Failure 400 {object} serializer.Response "参数错误" Example({"message": "参数错误"})
// @Router /users/activate/{token} [get]
func (v0) UserActivate(c *gin.Context) {
	userActive(c)

}

// UserLogout
// @Summary 用户登出接口
// @Description 用户登出账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Success 200 {object} serializer.Response "登出成功" Example({"message": "登出成功"})
// @Failure 400 {object} serializer.Response "参数错误" Example({"message": "参数错误"})
// @Router /users/session [delete]
func (v0) UserLogout(c *gin.Context) {
	userLogout(c)

}

// UserOauth2Login
// @Summary 用户第三方登录接口
// @Description 用户第三方登录账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param type path string true "用户第三方登录类型"
// @Param code path string true "用户第三方登录code"
// @Success 200 {object} serializer.Response "登录成功" Example({"message": "登录成功"})
// @Failure 400 {object} serializer.Response "参数错误" Example({"message": "参数错误"})
// @Router /users/oauth2/{type}/{code} [get]
func (v0) UserOauth2Login(c *gin.Context) {
	userOauth2Login(c)
}

// UserOauth2Bind
// @Summary 用户第三方绑定接口
// @Description 用户第三方绑定账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param type path string true "用户第三方绑定类型"
// @Param code path string true "用户第三方登录code"
// @Success 200 {object} serializer.Response "绑定成功" Example({"message": "绑定成功"})
// @Failure 400 {object} serializer.Response "参数错误" Example({"message": "参数错误"})
// @Router /users/oauth2/bind/{type}/{code} [get]
func (v0) UserOauth2Bind(c *gin.Context) {
	userOauth2Bind(c)
}
