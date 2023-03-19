package impl

import (
	"errors"
	"github.com/go-redis/redis"
	conf2 "github.com/huoxue1/tdlib/conf"
	log "github.com/sirupsen/logrus"
	"time"
)

type RedisClient struct {
	c *redis.Client
}

func (r *RedisClient) SetAdd(key, value string) error {
	return r.c.SAdd(key, value).Err()
}

func (r *RedisClient) SetIsMem(key string, member string) bool {
	isMember := r.c.SIsMember(key, member)
	if isMember.Err() != nil {
		return false
	}
	return isMember.Val()
}

func (r *RedisClient) SetDel(key string, member string) error {
	return r.c.SRem(key, member).Err()

}

func InitRedis(config *conf2.Config) (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Cache.Redis.Address,  // Redis 服务器地址和端口
		Password: config.Cache.Redis.Password, // Redis 认证密码，如果没有可以留空
		DB:       config.Cache.Redis.Db,       // Redis 数据库编号
	})
	if client == nil {
		log.Fatalln("初始化redis服务端失败")
		return nil, errors.New("初始化redis服务端失败")
	}
	return &RedisClient{c: client}, nil
}

func (r *RedisClient) Set(key, value string) error {
	return r.c.Set(key, value, 0).Err()
}

func (r *RedisClient) SetTtl(key, value string, ttl time.Duration) error {
	return r.c.Set(key, value, ttl).Err()
}

func (r *RedisClient) Get(key string) string {
	stringCmd := r.c.Get(key)
	if stringCmd.Err() != nil {
		log.Errorln(stringCmd.Err())
		return ""
	} else {
		result, err := stringCmd.Result()
		if err != nil {
			log.Errorln(err.Error())
			return ""
		}
		return result
	}
}

func (r *RedisClient) GetDefault(key string, defaultValue string) string {
	stringCmd := r.c.Get(key)
	if stringCmd.Err() != nil {
		log.Errorln(stringCmd.Err())
		return defaultValue
	} else {
		result, err := stringCmd.Result()
		if err != nil {
			log.Errorln(err.Error())
			return defaultValue
		}
		return result
	}
}

func (r *RedisClient) Delete(key string) error {
	return r.c.Del(key).Err()
}

func (r *RedisClient) ForEach(f func(key string, value string) bool) {
	scanCmd := r.c.Scan(0, "", 0)
	if scanCmd.Err() != nil {
		return
	}
	iterator := scanCmd.Iterator()
	for iterator.Next() {
		key := iterator.Val()
		value, err := r.c.Get(key).Result()
		if err != nil {
			continue
		}
		if !f(key, value) {
			break
		}
	}
}
