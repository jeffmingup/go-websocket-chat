package main

import (
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		// Dial or DialContext must be set. When both are set, DialContext takes precedence over Dial.
		Dial: func() (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
}

var (
	pool *redis.Pool
)

func init() {
	pool = newPool("localhost:6379")
}

func RedisDo(commandName string, args ...interface{}) (reply interface{}, err error) {
	if len(args) < 1 {
		return nil, errors.New("redis 命令缺失！")
	}
	args[0] = fmt.Sprintf("%s", args[0])

	conn := pool.Get()
	defer conn.Close()

	return conn.Do(commandName, args...)

}
