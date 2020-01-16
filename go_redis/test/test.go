package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func err_handler(err error) {
	fmt.Printf("err_handler, error:%s\n", err.Error())
	panic(err.Error())
}

func main() {
	// ----------------------------------------------------------------------
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	defer client.Close()

	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Printf("ping error[%s]\n", err.Error())
		err_handler(err)
	}
	fmt.Printf("ping result: %s\n", pong)

	fmt.Printf("----------------------------------------\n")

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Redis Dial:", "√")

	wg := &sync.WaitGroup{}
	expire := 2 * 60 * 60 * time.Second
	ct := 1000000
	for i := 0; i < ct; i++ {

		wg.Add(1)

		k := strconv.Itoa(i)
		v := rand.Int()
		err := client.Set("u:"+k, v, expire).Err()
		if err != nil {
			panic(err)
		}

		fmt.Println("redis set ", "u:"+k, v, "√")
		wg.Done()
	}
	wg.Wait()

	defer client.Close()

}
