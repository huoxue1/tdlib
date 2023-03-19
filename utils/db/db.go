package db

import (
	"github.com/huoxue1/tdlib/conf"
	"github.com/huoxue1/tdlib/utils/db/impl"
	"time"
)

type CacheClient interface {
	Set(key, value string) error
	SetTtl(key, value string, ttl time.Duration) error
	Get(key string) string
	GetDefault(key string, defaultValue string) string
	Delete(key string) error
	ForEach(func(key, value string) bool)

	SetAdd(key, value string) error
	SetIsMem(key string, member string) bool
	SetDel(key string, member string) error
}

var (
	c CacheClient
)

func GetCache() CacheClient {
	return c
}

func InitCache(config *conf.Config) error {
	if config.Cache.CacheType == "redis" {
		client, err := impl.InitRedis(config)
		if err != nil {
			return err
		}
		c = client
	} else {
		client, err := impl.InitNustdb(config)
		if err != nil {
			return err
		}
		c = client
	}
	return nil
}
