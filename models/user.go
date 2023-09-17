package models

import (
	"GoCloud/util"
	"gorm.io/gorm"
)

type User struct {
	// 表字段
	gorm.Model
	Email     string `gorm:"type:varchar(100);unique_index"`
	Nick      string `gorm:"size:50"`
	Password  string `json:"-"`
	Status    int
	GroupID   uint
	Storage   uint64
	TwoFactor string
	Avatar    string
	Options   string `json:"-" gorm:"size:4294967295"`
	Authn     string `gorm:"size:4294967295"`

	//// 关联模型
	//Group  Group  `gorm:"save_associations:false:false"`
	//Policy Policy `gorm:"PRELOAD:false,association_autoupdate:false"`
	//
	//// 数据库忽略字段
	//OptionsSerialized UserOption `gorm:"-"`
}

// todo 优化补充选项
// NewUser 返回一个新的空 User
func NewUser() User {
	//options := UserOption{}
	return User{
		//OptionsSerialized: options,
	}
}

// SetPassword 根据给定明文设定 User 的 Password 字段
func (user *User) SetPassword(password string) error {
	user.Password = util.NewCrypto(util.PasswordCrypto).Encrypt(password)
	return nil
}
