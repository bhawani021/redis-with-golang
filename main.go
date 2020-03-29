package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

func getRedisInstance() *redis.Client {
	redisClient :=  redis.NewClient(&redis.Options{
		// host:port address.
		Addr: "localhost:6379",
		// Maximum number of socket connections.
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: 100,
		// Maximum number of retries before giving up.
		// Default is to not retry failed commands.
		MaxRetries: 3,
		Password: "",
		DB: 0,
	})

	return redisClient
}

func SetValue(key string, value interface{}) (bool, error) {
	val, _ := json.Marshal(value)
	redisClient := getRedisInstance()
	err := redisClient.Set(key, string(val), 0).Err()
	return true, err
}

func GetValue(key string) (interface{}, error) {
	var val interface{}
	redisClient := getRedisInstance()
	result, err := redisClient.Get(key).Result()
	_ = json.Unmarshal([]byte(result), &val)
	return val, err
}

func main()  {
	// Set a new key
	set, err := SetValue("newKey", "newVal")
	if err == nil {
		fmt.Println(set)
	}

	// get key
	get, err := GetValue("newKey")
	if err == nil {
		fmt.Println(get)
	}
}