package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"sync"
	"time"
)

// redisStore redis存储驱动
type redisStore struct {
	client *redis.Client
	once   sync.Once
}
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
}

var redisStoreInstance = new(redisStore)

func newRedisStore(config RedisConfig) *redisStore {
	redisStoreInstance.once.Do(
		func() {
			redisStoreInstance.client = redis.NewClient(&redis.Options{
				Addr:     config.Addr,
				Password: config.Password,
				DB:       config.DB,
				PoolSize: config.PoolSize,
			})
		})
	return redisStoreInstance
}

// Get 键值对取值
func (r *redisStore) Get(key string) (any, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return val, errors.Wrapf(err, "redis get key: %s", key)
	}
	return val, nil
}

// Set 存储值
func (r *redisStore) Set(key string, value interface{}, ttl Second) error {
	ctx := context.Background()
	err := r.client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return errors.Wrapf(err, "redis set key: %s", key)
	}
	return nil
}

/* Hash Get*/

// HGet 获取hash值的某个字段
func (r *redisStore) HGet(key string, field string) (any, error) {
	ctx := context.Background()
	val, err := r.client.HGet(ctx, key, field).Result()

	if err != nil {
		return nil, errors.Wrapf(err, "redis hget key: %s", key)
	}
	return val, nil
}

// HMGet 获取hash值的多个字段
func (r *redisStore) HMGet(key string, fields []string) (map[string]interface{}, error) {
	ctx := context.Background()
	val, err := r.client.HMGet(ctx, key, fields...).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "redis hmget key: %s", key)
	}
	res := make(map[string]interface{})
	for i, v := range val {
		res[fields[i]] = v
	}
	return res, nil
}

// HGetAll 获取hash值的所有字段
func (r *redisStore) HGetAll(key string) (map[string]string, error) {
	ctx := context.Background()
	val, err := r.client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "redis hgetall key: %s", key)
	}
	return val, nil
}

/* Hash Set*/

// HSet 设置hash值的某个字段
func (r *redisStore) HSet(key string, field string, value interface{}, ttl Second) error {
	ctx := context.Background()
	err := r.client.HSet(ctx, key, field, value).Err()
	err = r.client.Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return errors.Wrapf(err, "redis hset key: %s", key)
	}
	return nil
}

// HMSet 设置hash值的所有字段
func (r *redisStore) HMSet(key string, values map[string]interface{}, ttl Second) error {
	ctx := context.Background()
	err := r.client.HMSet(ctx, key, values).Err()
	err = r.client.Expire(ctx, key, time.Duration(ttl)*time.Second).Err()
	if err != nil {
		return errors.Wrapf(err, "redis hmset key: %s", key)
	}
	return nil
}

// Gets 批量取值
func (r *redisStore) Gets(keys []string, prefix string) (map[string]interface{}, []string, error) {
	//TODO implement me
	panic("implement me")
	//ctx := context.Background()
	//// 构建带前缀的查询键
	//queryKeys := make([]string, len(keys))
	//for i, key := range keys {
	//	queryKeys[i] = prefix + key
	//}
	//// 执行 MGET 命令
	//result, err := r.client.MGet(ctx, queryKeys...).Result()
	//if err != nil {
	//	return nil, keys, errors.Wrapf(err, "redis gets keys: %s failed", keys)
	//}
	//
	//// 解码值并构建结果
	//res := make(map[string]any)
	//missed := make([]string, 0, len(keys))
	//
	//for i, value := range result {
	//	if value == nil {
	//		missed = append(missed, keys[i])
	//	} else {
	//		decoded, err := deserializer(value.([]byte))
	//		if err != nil || decoded == nil {
	//			missed = append(missed, keys[i])
	//		} else {
	//			res[keys[i]] = decoded
	//		}
	//	}
	//}
	//
	//return res, missed, nil
}

func (r *redisStore) Sets(values map[string]interface{}, prefix string) error {
	//TODO implement me
	panic("implement me")
}

func (r *redisStore) Delete(keys []string, prefix string) error {
	ctx := context.Background()
	queryKeys := make([]string, len(keys))
	for i, key := range keys {
		queryKeys[i] = prefix + key
	}
	err := r.client.Del(ctx, queryKeys...).Err()
	if err != nil {
		return errors.Wrapf(err, "redis delete keys: %s failed", keys)
	}
	return nil
}

// Persist 持久化
func (r *redisStore) Persist(path string) error {
	//TODO implement me
	panic("implement me")
}

// Restore 从文件中恢复
func (r *redisStore) Restore(path string) error {
	//TODO implement me
	panic("implement me")
}

// Clear 清空所有值
func (r *redisStore) Clear() error {
	err := r.client.FlushAll(context.Background()).Err()
	if err != nil {
		return errors.Wrapf(err, "redis clear failed")
	}
	return nil
}
