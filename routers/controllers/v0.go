package controllers

//进行多一层封装，以保证 api 的路径的修改方便，以及可以不用修改下层实现逻辑
import "github.com/gin-gonic/gin"

type v0 struct {
}

var V0 v0

//todo 补充更多信息

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
