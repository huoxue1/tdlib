package db

import (
	"github.com/go-redis/redis"
	"github.com/huoxue1/tdlib/conf"
	log "github.com/sirupsen/logrus"
)

var (
	client *redis.Client
)

func GetRedisClient() *redis.Client {
	return client
}

func InitRedis(config *conf.Config) {
	c := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    config.Redis.Address,
		OnConnect: func(conn *redis.Conn) error {

			log.Info("redis 连接成功")
			return nil
		},
		Password: config.Redis.Password,
	})
	if c == nil {
		log.Fatalln("redis连接失败")
	}
	client = c
}
