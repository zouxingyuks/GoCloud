package dao

import (
	"GoCloud/pkg/util"
	"GoCloud/service/cache"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"math/rand"
)

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
	DB().Where("email = ?", email).First(u)
	Cache().Set("email:"+email, u.UUID, cache.Second(600+rand.Intn(600)))
	data, err := util.StructToMapF1(u)
	if err != nil {
		return nil, errors.Wrap(err, "struct to map error")
	}
	//预防缓存雪崩，设置随机过期时间
	err = Cache().HMSet("uuid:"+u.UUID, data, cache.Second(600+rand.Intn(600)))
	if err != nil {
		return nil, errors.Wrap(err, "set cache error")
	}
	return u, nil
}

// 从缓存中查询用户信息
func findInCache(uuid string) (u *User, err error) {
	u = new(User)
	data, err := Cache().HGetAll("uuid:" + uuid)
	if err != nil {
		return nil, errors.Wrap(err, "get user from cache error")
	}
	err = mapstructure.Decode(data, u)
	if err != nil {
		return nil, errors.Wrap(err, "mapstructure decode error")
	}
	return u, nil
}
