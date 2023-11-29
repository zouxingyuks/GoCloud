package dao

import (
	"GoCloud/conf"
	"GoCloud/pkg/log"
	"fmt"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {
	conf.AddPath("../config")
	log.SetFilepath("../log/zap.log")
	fmt.Println(GetUser(WithUUID("ebbd708f-4ef3-5013-9a8a-bb28bd555e2e")))
	fmt.Println(GetUser(WithEmail("testtt@qq.com")))
	fmt.Println(GetUser(WithNickName("小可爱575bb7df85629bec0d8d7bde1172b59bb71fb6fb4413e1bc46860c520d5af81c")))
	time.Sleep(time.Second * 5)
}

func Test_checkEmailExist(t *testing.T) {
	conf.AddPath("../config")
	log.SetFilepath("../log/zap.log")
	fmt.Println(CheckEmailExist("zouxingyu@anubis.work"))
}
