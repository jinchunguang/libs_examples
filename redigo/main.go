package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
    "math/rand"
    "strconv"
    "sync"
    "time"
)

func main() {
    c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }

    _, err = c.Do("AUTH", "requirepass ", "")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    /*_, err = c.Do("SET", "mykey", "superWang")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }*/

    rand.Seed(time.Now().UnixNano())
    fmt.Println("Redis Dial:", "√")
    wg := &sync.WaitGroup{}
    ct := 1000000
    for i := 0; i < ct; i++ {

        wg.Add(1)

        k := strconv.Itoa(i)
        v := rand.Int()
        _, err = c.Do("SET", "uid:"+k, v)
        if err != nil {
            fmt.Println("redis set failed:", err)
        }
        fmt.Println("redis set ", "uid:"+k, v, "√")

        wg.Done()
    }
    wg.Wait()

    defer c.Close()
}
