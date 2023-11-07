package middleware

import (
	"GoCloud/pkg/log"
	"GoCloud/pkg/rbac"
	"GoCloud/pkg/util"
	"GoCloud/service/serializer"
	"github.com/gin-gonic/gin"
)

func RBAC(obj, act string) gin.HandlerFunc {
	return func(c *gin.Context) {
		entry := log.NewEntry("middleware.rbac")
		res := serializer.Response{}
		uuid := util.Session().Get(c, "user_id").(string)
		// 获取请求信息
		role, err := rbac.GetRolesForUser(uuid)
		if role == "" || err != nil {
			res = serializer.NewResponse(entry, 403, serializer.WithMsg("Permission denied"), serializer.WithErr(err))
			c.JSON(res.Code, res)
			c.Abort()
			return

		}
		// 检查策略
		ok, err, explains := rbac.Enforce(role, obj, act)
		if !ok || err != nil {
			res = serializer.NewResponse(entry, 403, serializer.WithMsg("Permission denied"), serializer.WithErr(err), serializer.WithData(explains))
			c.JSON(res.Code, res)
			c.Abort()
			return
		}
		c.Next()
	}
}
