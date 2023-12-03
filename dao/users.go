package dao

import (
	"GoCloud/pkg/util"
	"GoCloud/service/cache"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/rand"
)

// 从缓存中查询用户信息
func findUserInCache(uuid string) (u *User, err error) {
	u = new(User)
	data, err := Cache().HGetAll("uuid:" + uuid)
	if err != nil {
		return nil, errors.Wrap(err, "get user from cache error")
	}
	if len(data) == 0 {
		return nil, errors.New("user not found")
	}
	err = mapstructure.Decode(data, u)
	if err != nil {
		return nil, errors.Wrap(err, "mapstructure decode error")
	}
	return u, nil
}

// storeUserInCache 将用户信息存入缓存
func storeUserInCache(u *User) (err error) {
	//预防缓存雪崩，设置随机过期时间
	ttl := cache.Second(600 + rand.Intn(600)) // 10min~20min
	// email 与 uuid 的映射
	err = Cache().Set("email:"+u.Email, u.UUID, ttl)
	if err != nil {
		return errors.Wrap(err, "set cache error")
	}
	data, err := util.StructToMapF1(u)
	if err != nil {
		return errors.Wrap(err, "struct to map error")
	}
	// uuid 与 user 的映射
	err = Cache().HMSet("uuid:"+u.UUID, data, ttl)
	if err != nil {
		return errors.Wrap(err, "set cache error")
	}
	return nil
}

// updateUserInCache 更新缓存中的用户信息
// 根据 uuid 定位旧的 User ，然后将新的 User 的信息更新到旧的 User 中，只更新非空字段
func updateUserInCache(uuid string, newUser *User) (err error) {
	//预防缓存雪崩，设置随机过期时间
	ttl := cache.Second(600 + rand.Intn(600)) // 10min~20min
	data, err := util.StructToMapF1(newUser)
	delete(data, "uuid")
	for k, v := range data {
		if v == "" {
			delete(data, k)
		}
	}
	//更新
	err = Cache().HMSet("uuid:"+uuid, data, ttl)
	if err != nil {
		return errors.Wrap(err, "set cache error")
	}
	return nil

}

// refreshUserInCache 刷新缓存中的用户信息
func refreshUserInCache(uuid string) (u *User, err error) {
	u = new(User)
	err = DB().Where("uuid = ?", uuid).First(u).Error
	if err != nil {
		return u, errors.Wrap(err, "get user from db error")
	}
	err = storeUserInCache(u)
	if err != nil {
		return nil, errors.Wrap(err, "set cache error")
	}
	return u, nil
}

// GetUserByEmail 根据邮箱获取用户信息
func GetUserByEmail(email string) (u *User, err error) {
	u = new(User)
	//1. 先在缓存中查找
	t, err := Cache().Get("email:" + email)
	u.UUID = t.(string)
	if err == nil {
		return findUserInCache(u.UUID)
	}
	//2. 如果缓存中没有，就去数据库中查找
	u = new(User)
	err = DB().Where("email = ?", email).First(u).Error
	if err != nil {
		return u, errors.Wrap(err, "get user from db error")
	}
	err = storeUserInCache(u)
	if err != nil {
		return nil, errors.Wrap(err, "set cache error")
	}
	return u, nil
}

// GetUserByUUID 根据UUID获取用户信息
func GetUserByUUID(uuid string) (*User, error) {
	if uuid == "" {
		return nil, errors.New("uuid is empty")
	}
	//先在缓存中查找
	u, err := findUserInCache(uuid)
	if err == nil && u != nil {
		return u, nil
	}
	//如果缓存中没有再去数据库中查找
	u = new(User)
	err = DB().Where("uuid = ?", uuid).First(u).Error
	if err != nil {
		return u, errors.Wrap(err, "get user from db error")
	}
	err = storeUserInCache(u)
	if err != nil {
		return nil, errors.Wrap(err, "set cache error")
	}
	return u, nil
}

// SetUser 设置用户信息
func SetUser(uuid string, opts ...UserOption) (u *User, err error) {
	//设置目标用户
	o := new(User)
	for _, opt := range opts {
		opt.apply(o)
	}
	//先在缓存中查找，更新一次缓存，用于保存用户信息
	u, err = refreshUserInCache(uuid)
	err = DB().Where("uuid = ?", uuid).Session(&gorm.Session{NewDB: false}).Updates(o).Error
	if err != nil {
		//如果更新失败，就把缓存中的用户信息还原
		err = DB().Where("uuid = ?", uuid).Session(&gorm.Session{NewDB: false}).Updates(u).Error
		return nil, errors.Wrap(err, "set user error")
	}
	//更新缓存
	u, err = refreshUserInCache(uuid)
	return u, err
}
