package dao

import (
	"GoCloud/conf"
	"GoCloud/pkg/log"
	"testing"
	"time"
)

func TestGetUserByEmail(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	time.Sleep(time.Second * 5)
	u, err := GetUserByEmail("1308345487@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func Test_findUserInCache(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	time.Sleep(time.Second * 5)
	u, err := findUserInCache("8ae508bc-1e19-5795-93c9-860440c393ac")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestGetUserByUUID(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	time.Sleep(time.Second * 5)
	u, err := GetUserByUUID("8ae508bc-1e19-5795-93c9-860440c393ac")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}