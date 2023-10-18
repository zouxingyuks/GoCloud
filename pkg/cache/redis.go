package cache

import (
	"GoCloud/pkg/conf"
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

var redisStoreInstance = new(redisStore)

func RedisStore() *redisStore {
	redisStoreInstance.once.Do(
		func() {
			redisStoreInstance.client = redis.NewClient(&redis.Options{
				Addr:     conf.RedisConfig().Server,
				Password: conf.RedisConfig().Password,
				DB:       conf.RedisConfig().DB,
				PoolSize: conf.RedisConfig().PoolSize,
			})
		})
	return redisStoreInstance
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

// Get 取值
func (r *redisStore) Get(key string) (interface{}, error) {
	ctx := context.Background()
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return val, errors.Wrapf(err, "redis get key: %s", key)
	}
	return val, nil
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
