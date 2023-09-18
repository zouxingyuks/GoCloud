package models

import (
	"GoCloud/util/crypto"
	"errors"
	"gorm.io/gorm"
	"regexp"
)

const (
	// UserActive 账户正常状态
	UserActive = iota
	// NotActivicated 未激活
	NotActivicated
	// Baned 被封禁
	Baned
	// OveruseBaned 超额使用被封禁
	OveruseBaned
)

type User struct {
	// 表字段
	gorm.Model
	UserName  string `gorm:"type:varchar(100);unique_index"`
	Email     string `gorm:"type:varchar(100);unique_index"`
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
	ciphertext, err := crypto.NewCrypto(crypto.PasswordCrypto).Encrypt([]byte(password))
	if err != nil {
		return err
	}
	user.Password = ciphertext
	return nil
}

// SetUserName 用于设置用户的用户名
func (user *User) SetUserName(username string) error {
	// 长度要求：用户名应至少包含3到20个字符。
	if len(username) < 3 || len(username) > 20 {
		return errors.New("用户名长度必须在3到20个字符之间")
	}

	// 字符允许性：只允许使用字母（大小写都可以）、数字和下划线作为用户名的一部分。
	regex := "^[a-zA-Z0-9_]+$"
	match, _ := regexp.MatchString(regex, username)
	if !match {
		return errors.New("用户名只能包含字母、数字和下划线")
	}
	//todo  检查是否存在同名用户

	// 如果所有规则都通过，则设置用户名
	user.UserName = username
	return nil
}
