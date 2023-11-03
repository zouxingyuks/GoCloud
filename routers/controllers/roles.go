package controllers

import (
	"GoCloud/pkg/log"
	"GoCloud/pkg/serializer"
	"GoCloud/service/rbac"
	"github.com/gin-gonic/gin"
)

// AssignRolesToUser 给用户分配角色的接口。
// @Summary 给用户分配角色
// @Description 为用户分配一个或多个角色。
// @Tags 用户管理
// @Accept application/json
// @Produce application/json
// @Param userId path string true "用户ID"
// @Param roles body RolesParam true "角色标识符列表"
// @Success 200 {object} Response "分配成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 404 {object} ErrorResponse "用户或角色不存在"
// @Failure 500 {object} ErrorResponse "服务器错误"
// @Router /users/{userId}/roles [post]
func (v0) AssignRolesToUser(c *gin.Context) {
	// 实现给用户分配角色的逻辑
	entry := log.NewEntry("controller.user.roles")
	service := rbac.NewService()
	res := serializer.Response{}
	if err := c.ShouldBindJSON(&service); err == nil {
		// 注册用户
		res = service.AssignRolesToUser(c)
	} else {
		// 参数错误
		res = serializer.NewResponse(entry, 400, serializer.WithMsg("参数错误"), serializer.WithErr(err))
	}
	c.JSON(res.Code, res)
}
