package dao

import (
	"GoCloud/conf"
	"GoCloud/pkg/log"
	"testing"
)

func TestGetUserByEmail(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	u, err := GetUserByEmail("1308345487@qq.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func Test_findUserInCache(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	u, err := findUserInCache("8ae508bc-1e19-5795-93c9-860440c393ac")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func TestGetUserByUUID(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	u, err := GetUserByUUID("8ae508bc-1e19-5795-93c9-860440c393ac")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}

func Test_updateUserInCache(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	err := updateUserInCache("1182bd1f-9d5d-55f9-bb2f-7104adfbb513", &User{
		NickName: "test",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_refreshUserInCache(t *testing.T) {
	conf.AddPath("..")
	log.SetFilepath("../log/zap.log")
	u, err := refreshUserInCache("1182bd1f-9d5d-55f9-bb2f-7104adfbb513")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u)
}
