package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

type redisSingleObj struct {
	Redis_host string
	Redis_port uint16
	Redis_auth string
	Database   int
	Db         *redis.Client
}

func (r *redisSingleObj) InitSingleRedis() (err error) {
	redisAddr := fmt.Sprintf("%s:%d", r.Redis_host, r.Redis_port)
	r.Db = redis.NewClient(&redis.Options{
		Addr:        redisAddr,
		Password:    r.Redis_auth,
		DB:          r.Database,
		IdleTimeout: 300,
		PoolSize:    100,
	})
	fmt.Printf("connecting redis:%v\n", redisAddr)

	res, err := r.Db.Ping().Result()
	if err != nil {
		fmt.Printf("Connect to redis error:%v\n", err)
		return err
	} else {
		fmt.Printf("Conncet to redis success:%v\n", res)
		return nil
	}

}

func main() {
	conn := &redisSingleObj{
		Redis_host: "127.0.0.1",
		Redis_port: 6379,
		Redis_auth: "3953",
	}
	err := conn.InitSingleRedis()
	if err != nil {
		panic(err)
	}
	defer conn.Db.Close()
}
