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
