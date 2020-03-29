// Storing composite values
package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
)

const keyToAdd  = "key009"

type Employee struct {
	Name    string  `json:"name"`
	Age     int     `json:"age"`
	Salary  string	`json:"salary"`
}

func main()  {
	// Create redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		PoolSize: 100,
		MaxRetries: 3,
		Password: "",
		DB: 0,
	})

	// Employee structure
	emp := Employee{
		Name: "Bhawani Shanker",
		Age: 36,
		Salary: "Test",
	}

	// Marshalling
	jsonObj, err := json.Marshal(emp)
	if err != nil {
		panic(err)
	}

	// Set value in Redis DB
	err = redisClient.Set(keyToAdd, jsonObj, 0).Err()
	if err != nil {
		panic(err)
	}

	// Get value from Redis DB
	res, err := redisClient.Get(keyToAdd).Result()
	if err != nil {
		panic(err)
	}

	// Unmarshalling result
	_ = json.Unmarshal([]byte(res), &emp)
	
	fmt.Println("result:", emp)
}


