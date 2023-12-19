package space

import "testing"

func TestDeleteSpace(t *testing.T) {
	err := DeleteSpace("testtt")
	if err != nil {
		t.Error(err)
	}

}

func TestNewSpace(t *testing.T) {
	err := NewSpace("local")
	if err != nil {
		t.Error(err)
	}
}

func TestExistSpace(t *testing.T) {
	exist, err := ExistSpace("local")
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Error("集合不存在")
	}
}
