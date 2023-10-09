package model

import "sync"

type Mail struct {
	Type int
	Once sync.Once `mapstructure:"-"`
	//同时最大发送数量
	Size int
	Smtp SMTP
}

type SMTP struct {
	Host string
	Once sync.Once `mapstructure:"-"`
	//端口
	Port int
	//用户名
	User string
	//密码
	Password string
	//发送者名
	Name string
	//发送者地址
	Address string
	//回复地址
	ReplyTo string
	//是否启用加密
	Encryption bool
	//SMTP 连接保留时长
	Keepalive int
}
