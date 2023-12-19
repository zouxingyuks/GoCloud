package space

import "testing"

func TestDeleteCollection(t *testing.T) {
	err := DeleteCollection("testtt")
	if err != nil {
		t.Error(err)
	}

}

func TestNewCollection(t *testing.T) {
	err := NewCollection("local")
	if err != nil {
		t.Error(err)
	}
}

func TestExistCollection(t *testing.T) {
	exist, err := ExistCollection("local")
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Error("集合不存在")
	}
}
