package dao

import (
	"GoCloud/conf"
	"GoCloud/pkg/log"
	"fmt"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	time.Sleep(time.Second * 5)
	fmt.Println(GetUserByUUID("8ae508bc-1e19-5795-93c9-860440c393ac"))
}

func Test_checkEmailExist(t *testing.T) {
	conf.AddPath("../config")
	log.SetFilepath("../log/zap.log")
	fmt.Println(GetUserByEmail("zouxingyu@anubis.work"))
}
