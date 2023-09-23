package controllers

import (
	"GoCloud/pkg/serializer"
	user2 "GoCloud/service/user"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 所以的参数过滤均再此处过滤
// 包括但不限于传入参数过滤、返回信息过滤
// todo 参数过滤，
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
		c.JSON(res.Code, res)

	} else {
		// 参数错误
		res = serializer.Err(
			http.StatusBadRequest,
			"参数错误",
			err)
		//请求失败，返回400
		c.JSON(400, res)
	}
}
func userLogin(c *gin.Context) {
	user := user2.NewUserService()
	res := serializer.Response{}
	if err := c.ShouldBindJSON(&user); err == nil {
		// 注册用户
		res = user.Login(c)
		//返回基本用户数据
		c.JSON(res.Code, res)

	} else {
		// 参数错误
		res = serializer.Err(
			http.StatusBadRequest,
			"参数错误",
			err)
		//请求失败，返回400
		c.JSON(400, res)
	}
}
