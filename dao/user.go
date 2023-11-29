package dao

import (
	"GoCloud/pkg/log"
	"GoCloud/pkg/uuid"
	"GoCloud/service/cache"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"math/rand"

	"gorm.io/gorm"
	"regexp"
)

const (
	// UserActive 账户正常状态
	UserActive = iota
	// UserNotActivated 未激活
	UserNotActivated
	// UserBaned 被封禁
	UserBaned
	// OveruseBaned 超额使用被封禁
	OveruseBaned
)

// User 此处UUID，NickName，Email都是唯一键，但是只有UUID是主键，且UUID不可修改
type User struct {
	// 表字段
	gorm.Model `json:"-" json:"gorm.Model"`
	UUID       string `gorm:"type:varchar(100);unique_index" json:"UUID,omitempty"`
	NickName   string `gorm:"type:varchar(100);unique_index" json:"nickName,omitempty"`
	Email      string `gorm:"type:varchar(100);unique_index" json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Status     int    `json:"status,omitempty"`
	Storage    uint64 `json:"storage,omitempty"`
	TwoFactor  string `json:"twoFactor,omitempty"`
	Avatar     string `json:"avatar,omitempty"`
	//// 关联模型
	//Group  Group  `gorm:"save_associations:false:false"`
	//Policy Policy `gorm:"PRELOAD:false,association_autoupdate:false"`
	//
	//// 数据库忽略字段
	//OptionsSerialized UserOption `gorm:"-"`
}

// NewUser 返回一个新的空 User
func NewUser(opts ...UserOption) *User {
	o := new(User)
	for _, opt := range opts {
		opt.apply(o)
	}
	//todo 唯一键的选择
	o.UUID = uuid.New(uuid.WitName(o.Email)).Get()
	o.NickName = "user" + o.UUID
	return o
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

// UserOption 定义一个接口类型
type UserOption interface {
	apply(*User)
}

// funcOption 定义funcOption类型，实现 IOption 接口
type funcOption struct {
	f func(*User)
}

func (fo funcOption) apply(u *User) {
	fo.f(u)
}

func newFuncOption(f func(option *User)) UserOption {
	return &funcOption{
		f: f,
	}
}

// WithEmail 定义一个函数，用于设置 Email
func WithEmail(email string) UserOption {
	return newFuncOption(func(o *User) {
		o.Email = email
	})
}

// WithStatus 定义一个函数，用于设置 Status
func WithStatus(status int) UserOption {
	return newFuncOption(func(o *User) {
		o.Status = status
	})
}

// WithNickName 定义一个函数，用于设置 NickName
func WithNickName(nickName string) UserOption {
	return newFuncOption(func(o *User) {
		o.NickName = nickName
	})
}

// WithUUID 定义一个函数，用于设置 UserID
func WithUUID(uuid string) UserOption {
	return newFuncOption(func(o *User) {
		o.UUID = uuid
	})
}

// GetUser 获取用户信息
func GetUser(opts ...UserOption) (User, error) {
	var (
		u   = new(User)
		err error
		o   = new(User)
	)

	for _, opt := range opts {
		opt.apply(o)
	}
	if o.UUID != "" {
		//先在缓存中查找
		u, ok := findInCache(o)
		if ok {
			return *u, nil
		}
	}
	u, err = findUserFromDB(o)
	return *u, err
}

// CreateUser 创建用户
func CreateUser(u *User) (err error) {
	err = DB().Session(&gorm.Session{NewDB: false}).Create(u).Error
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

// todo 这部分是否需要额外加锁

// SetUser 设置用户信息
func SetUser(opts ...UserOption) (u User, err error) {
	o := new(User)
	//此处是先从数据库中获取用户信息
	*o, err = GetUser(opts...)
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

// 在缓存中查询
// 此处传地址主要是为了节省内存
func findInCache(user *User) (u *User, ok bool) {
	entry := log.NewEntry("dao")
	//用uuid查找user
	result, err := Cache().Get("user:" + user.UUID)
	if err == nil {
		err = json.Unmarshal([]byte(result.(string)), &u)
		if err != nil {
			entry.Error(fmt.Sprintf("%s found in cache but unmarshal error", user.UUID))
			return nil, false
		}
		return u, true
	}
	entry.Debug(fmt.Sprintf("%s  not found in cache", user.UUID), log.Field{Key: "err", Value: err.Error()})
	return nil, false
}

// 从数据库中查询
// 此处传地址主要是为了节省内存
func findUserFromDB(user *User) (u *User, err error) {
	entry := log.NewEntry("dao")
	u = new(User)
	//缓存中没有再在数据库中查找
	err = DB().Where(&user).First(u).Error

	entry.Debug("get user from db", log.Field{
		Key:   "user",
		Value: user,
	})
	if err != nil {
		return u, errors.Wrap(err, "get user from db error")
	}
	//存入缓存
	//预防缓存雪崩，设置随机过期时间
	ttl := cache.Second(600 + rand.Intn(600)) // 10min~20min

	// email 与 uuid 的映射
	err = Cache().Set("email:"+user.Email, user.UUID, ttl)
	if err != nil {
		return u, errors.Wrap(err, "set cache error")
	}

	// uuid 与 user 的映射
	jsonData, err := json.Marshal(user)
	if err != nil {
		return u, errors.Wrap(err, "json marshal error")
	}

	err = Cache().Set("user:"+user.UUID, jsonData, ttl)
	if err != nil {
		return u, errors.Wrap(err, "set cache error")
	}
	return u, nil
}

func checkPassword(email string, password string) bool {
	//todo
	uuid, err := Cache().Get("email:" + email)
	if err != nil {
		return false
	}
	Cache().HGet("user:"+uuid.(string), "password")
	return true
}
