package main

import (
    "log"
    "time"
)

func main() {
    var i int

    for i = 0; i < 10; i ++ {
        go func(i int) {
            log.Println(i)
        }(i)
    }

    time.Sleep(2 * time.Second)
}
