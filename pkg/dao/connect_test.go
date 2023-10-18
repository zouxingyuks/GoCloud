package dao

import (
	"GoCloud/pkg/conf"
	"testing"
)

func TestDB(t *testing.T) {
	conf.AddPath("../../config")
	DB()
}
