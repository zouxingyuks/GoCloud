package controllers

//进行多一层封装，以保证 api 的路径的修改方便，以及可以不用修改下层实现逻辑
import "github.com/gin-gonic/gin"

type V0 struct {
}

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
