package dao

import (
	"rateLimitMiddleware/conf"

	"github.com/go-redis/redis"
)

var (
	// DB ...
	DB *redis.Client
)

// ConnectDataBase ...
func ConnectDataBase() error {
	// DB = redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })
	// fmt.Println("my log:" + DB.ClientGetName().String())
	opt, err := redis.ParseURL(conf.DatabaseAddr + conf.DatabaseName)
	if err != nil {
		return err
	}

	DB = redis.NewClient(opt)
	return nil
}
