package main

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func main() {
	l := rate.NewLimiter(1, 3)
	for {
		r := l.ReserveN(time.Now(), 3)
		time.Sleep(r.Delay())
		fmt.Println(time.Now().Format("2016-01-02 15:04:05.000"))
	}
}
