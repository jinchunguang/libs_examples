package main

import (
    "container/ring"
    "fmt"
)

func main() {
    ring := ring.New(10)

    for i := 1; i <= 10; i++ {
        ring.Value = i
        ring = ring.Next()
    }

    // 计算 1+2+3
    s := 0
    ring.Do(func(p interface{}){
        s += p.(int)

        fmt.Println(p.(int))
    })
    fmt.Println("sum is", s)
}