package configs

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	RedisClient1 *Redis
	RedisClient2 *Redis
)

func InitRedis() {

	for name, conf := range Resource.Redis {
		//MysqlClients[name],_ = initMysqlClient(conf)
		switch name {
		case "redis1":
			RedisClient1, _ = initRedisClient(conf)
		case "redis2":
			RedisClient2, _ = initRedisClient(conf)
		}
	}
}

type Redis struct {
	pool       *redis.Pool
	Service    string
	RemoteAddr string
}

func initRedisClient(conf RedisConf) (*Redis, error) {
	p := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				"tcp",
				conf.Addr,
				redis.DialPassword(conf.Password),
				redis.DialConnectTimeout(conf.ConnTimeOut),
				redis.DialReadTimeout(conf.ReadeTimeOut),
				redis.DialWriteTimeout(conf.WriteTimeOut),
			)
			if err != nil {
				return nil, err
			}
			return conn, err
		},
		DialContext: nil,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:         conf.MaxIdle,
		MaxActive:       conf.MaxActive,
		IdleTimeout:     conf.IdleTimeout,
		Wait:            true,
		MaxConnLifetime: conf.MaxConnLifetime,
	}
	c := &Redis{
		pool:       p,
		Service:    conf.Service,
		RemoteAddr: conf.Addr,
	}
	return c, nil
}
