package dao

import (
	"GoCloud/pkg/crypto"
	"GoCloud/pkg/log"
	"GoCloud/pkg/uuid"
	"github.com/pkg/errors"

	"gorm.io/gorm"
	"regexp"
)

const (
	// UserActive 账户正常状态
	UserActive = iota
	// UserNotActivated 未激活
	UserNotActivated
	// Baned 被封禁
	Baned
	// OveruseBaned 超额使用被封禁
	OveruseBaned
)

type User struct {
	// 表字段
	gorm.Model
	UUID      string `gorm:"type:varchar(100);unique_index"`
	NickName  string `gorm:"type:varchar(100);unique_index"`
	Email     string `gorm:"type:varchar(100);unique_index"`
	Password  string `json:"-"`
	Status    int
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

// NewUser 返回一个新的空 User
func NewUser(opts ...UserOption) User {
	o := new(userOption)
	for _, opt := range opts {
		opt.apply(o)
	}
	//todo 唯一键的选择
	o.User.UUID = uuid.New(uuid.WitName(o.User.Email)).Get()
	return o.User
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

// todo 怎么用这个
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
	user.NickName = username
	return nil
}

// 类型名称及字段的首字母小写（包内私有）
type userOption struct {
	User
}

// UserOption 定义一个接口类型
type UserOption interface {
	apply(*userOption)
}

// funcOption 定义funcOption类型，实现 IOption 接口
type funcOption struct {
	f func(*userOption)
}

func (fo funcOption) apply(o *userOption) {
	fo.f(o)
}

func newFuncOption(f func(option *userOption)) UserOption {
	return &funcOption{
		f: f,
	}
}

// WithEmail 定义一个函数，用于设置 Email
func WithEmail(email string) UserOption {
	return newFuncOption(func(o *userOption) {
		o.Email = email
	})
}

// WithStatus 定义一个函数，用于设置 Status
func WithStatus(status int) UserOption {
	return newFuncOption(func(o *userOption) {
		o.Status = status
	})
}

// WithNickName 定义一个函数，用于设置 NickName
func WithNickName(nickName string) UserOption {
	return newFuncOption(func(o *userOption) {
		o.NickName = nickName
	})
}

// WithUserUUID 定义一个函数，用于设置 UserID
func WithUserUUID(uuid string) UserOption {
	return newFuncOption(func(o *userOption) {
		o.UUID = uuid
	})
}

// GetUser 获取用户信息
func GetUser(opts ...UserOption) (u User, err error) {

	o := new(userOption)
	for _, opt := range opts {
		opt.apply(o)
	}
	err = DB().Where(o.User).First(&u).Error
	log.NewEntry("dao").Debug("GetUser", log.Field{
		Key:   "user",
		Value: u,
	}, log.Field{
		Key:   "err",
		Value: err,
	})
	return u, err
}

// CreateUser 创建用户
func CreateUser(u *User) (err error) {
	DB().Session(&gorm.Session{NewDB: false})

	err = DB().Create(u).Error
	log.NewEntry("dao").Debug("CreateUser", log.Field{
		Key:   "user",
		Value: u,
	},
		log.Field{
			Key:   "err",
			Value: err,
		})
	return err
}

// todo 如何设置密码呢
// todo 这部分是否需要额外加锁

// SetUser 设置用户信息
func SetUser(opts ...UserOption) (u User, err error) {
	o := new(userOption)
	//此处是先从数据库中获取用户信息
	o.User, err = GetUser(opts...)
	if err != nil {
		return u, err
	}
	//进行修改
	for _, opt := range opts {
		opt.apply(o)
	}
	//存回数据库
	err = DB().Save(&u).Error
	return u, err
}
