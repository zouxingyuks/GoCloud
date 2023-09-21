package controllers

import (
	"GoCloud/pkg/serializer"
	user2 "GoCloud/service/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// userRegister 用户注册接口
func userRegister(c *gin.Context) {
	user := user2.NewUserService()
	//todo 在参数设置这里尝试进行优化

	res := serializer.Response{}

	// 解析并验证用户注册信息
	if err := c.ShouldBindJSON(&user); err == nil {
		//todo 删除测试
		fmt.Println(user)
		// 注册用户
		res = user.Register()
	} else {
		// 参数错误
		res = serializer.Response{
			Code: http.StatusBadRequest,
			Msg:  "参数错误",
		}
	}

	c.JSON(res.Code, res)
}
