package active

import (
	"GoCloud/dao"
	"GoCloud/pkg/log"
	"GoCloud/pkg/serializer"
	"github.com/gin-gonic/gin"
	"time"
)

const activeAction = "active"

var exp = time.Hour

// UserActivate
// @Summary 用户激活接口
// @Description 用户激活账户
// @Tags User
// @Accept application/json
// @Produce application/json
// @Param token path string true "用户激活token"
// @Success 200 {object} serializer.Response "激活成功"
// @Failure 400 {object} serializer.Response "参数错误"
// @Failure 500 {object} serializer.Response "服务异常"
// @Router /users/activate/{token} [get]
func UserActivate(c *gin.Context) {
	entry := log.NewEntry("controller.user.active")

	tokenStr := c.Param("token")
	// 1.参数校验
	if tokenStr == "" {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg("参数错误"))
		c.JSON(res.Code, res)
		return
	}
	// 2.解析token
	sub, err := Parse(tokenStr, activeAction)
	if err != nil {
		res := serializer.NewResponse(entry, 400, serializer.WithMsg("无效的token"), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	// 3. 激活用户
	uuid := sub.(string)
	_, err = dao.SetUser(uuid, dao.WithStatus(dao.UserActive))
	if err != nil {
		res := serializer.NewResponse(entry, 500, serializer.WithMsg("服务异常"), serializer.WithErr(err))
		c.JSON(res.Code, res)
		return
	}
	// 链接使用的日志记录
	res := serializer.NewResponse(entry, 200, serializer.WithMsg("激活成功"))
	c.JSON(res.Code, res)
	return
}
