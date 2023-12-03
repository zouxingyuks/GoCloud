package dao

import (
	"GoCloud/pkg/util"
	"GoCloud/service/cache"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"math/rand"
)

// 从缓存中查询用户信息
func findInCache(uuid string) (u *User, err error) {
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

// storeInCache 将用户信息存入缓存
func storeInCache(u *User) (err error) {
	//预防缓存雪崩，设置随机过期时间
	ttl := cache.Second(600 + rand.Intn(600)) // 10min~20min
	// email 与 uuid 的映射
	Cache().Set("email:"+u.Email, u.UUID, ttl)
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

// GetUserByEmail 根据邮箱获取用户信息
func GetUserByEmail(email string) (u *User, err error) {
	u = new(User)
	//1. 先在缓存中查找
	t, err := Cache().Get("email:" + email)
	u.UUID = t.(string)
	if err == nil {
		return findInCache(u.UUID)
	}
	//2. 如果缓存中没有，就去数据库中查找
	u = new(User)
	err = DB().Where("email = ?", email).First(u).Error
	if err != nil {
		return u, errors.Wrap(err, "get user from db error")
	}
	err = storeInCache(u)
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
	u, err := findInCache(uuid)
	if err == nil && u != nil {
		return u, nil
	}
	//如果缓存中没有再去数据库中查找
	u = new(User)
	err = DB().Where("uuid = ?", uuid).First(u).Error
	if err != nil {
		return u, errors.Wrap(err, "get user from db error")
	}
	err = storeInCache(u)
	if err != nil {
		return nil, errors.Wrap(err, "set cache error")
	}
	return u, nil
}
