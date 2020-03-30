package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

const KeyToAdd  = "test010"

func main()  {

	redisClient := redis.NewClient(&redis.Options{
		// host:port address
		Addr: "localhost:6379",
		// Maximum number of socket connections
		// Default is 10 connections per every CPU as reported by runtime.NumCPU.
		PoolSize: 100,
		// Maximum number of retries before giving up
		// Default is to not retry failed commands
		MaxRetries: 2,
		Password: "",
		DB: 0,
	})

	// Set value
	var val string
	val = "my testing value"
	jsonObj, _ := json.Marshal(val)
	_ = redisClient.Set(KeyToAdd, jsonObj, 0).Err()

	// Get value
	var resultData string
	result, _ := redisClient.Get(KeyToAdd).Result()
	_ = json.Unmarshal([]byte(result), &resultData)
	fmt.Println(fmt.Println("resultData =", resultData)) // resultData = my testing value

	// Delete key
	_ = redisClient.Del(KeyToAdd).Err()

	// Set value with ttl of 10 secs
	_ = redisClient.Set(KeyToAdd, jsonObj, time.Duration(10)*time.Second).Err()

	// PUT sleep for 10 secs
	time.Sleep(time.Second * 10)

	// GET value
	var resultData1 string
	result, _ = redisClient.Get(KeyToAdd).Result()
	_ = json.Unmarshal([]byte(result), &resultData1)
	fmt.Println("resultData1 =", resultData1) // resultData1 =
}


