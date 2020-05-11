package gredis

import (
	"encoding/json"
	"time"

	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
)

var RedisConn *redis.Pool

func Setup() error {
	RedisConn = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			dial, err := redis.Dial("tcp", viper.GetString("redis.host"))
			if err != nil {
				return nil, err
			}
			if viper.Get("redis.password") != nil {
				if _, err = dial.Do("AUTH", viper.GetString("redis.password")); err != nil {
					dial.Close()
					return nil, err
				}
			}
			return dial, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
		MaxIdle:     30,
		MaxActive:   30,
		IdleTimeout: 200,
	}
	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func Exist(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	b, err := redis.Bool(conn.Do("exist", key))
	if err != nil {
		return false
	}
	return b
}
