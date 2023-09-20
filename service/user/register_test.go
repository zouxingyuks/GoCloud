package user

import (
	"GoCloud/models"
	"GoCloud/pkg/conf"
	"fmt"
	"testing"
)

func Test_defaultUserName(t *testing.T) {
	email := "1111@qq.com"
	fmt.Println(defaultUserName(email))
}

func TestParam_Register(t *testing.T) {
	//测试设置
	conf.TestMode()
	models.InitDao()
	param := Param{
		Email:    "test@qq.com",
		Password: "123456",
	}
	fmt.Println(param.Register())
}
