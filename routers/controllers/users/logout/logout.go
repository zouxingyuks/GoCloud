package logout

import "C"
import (
	"GoCloud/pkg/log"
	"GoCloud/pkg/serializer"
	"GoCloud/pkg/util"
	"github.com/gin-gonic/gin"
)

// Logout 用户登出
// @Summary 用户登出
// @Description 用户登出
// @Tags User
// @Accept application/json
// @Produce application/json
// @Success 200 {object} string "登出成功"
// @Failure 401 {object} string "身份验证失败"
// @Router /users/session [delete]
func Logout(c *gin.Context) {
	entry := log.NewEntry("controller.user.logout")

	// 1. 验证用户身份
	uuid := util.Session().Get(c, "uuid")
	if uuid == nil {
		// 1.1 身份验证失败
		res := serializer.NewResponse(entry, 401, serializer.WithMsg("身份验证失败"))
		c.JSON(res.Code, res)
		return
	}
	// 2. 终止用户会话
	// 2.1 删除redis中的会话信息
	util.Session().Delete(c, "uuid")

	// 3 记录登出日志
	res := serializer.NewResponse(entry, 200, serializer.WithMsg(uuid.(string)+"登出成功"))
	c.JSON(res.Code, res)
}
