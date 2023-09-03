package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func connectionToRedis() (redis.Conn, error) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis conn failed", err)
		return nil, err
	}
	return c, err

}
