package main

import (
	"fmt"
	"time"
)

// 一次性任务
func main() {
	// timer1()

	timer2()
}

// 重置时间
func timer2() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	time.Sleep(1 * time.Second)
	timer.Reset(6 * time.Second)
	x, _ := <-timer.C
	fmt.Println(x)

}

// 停止计时器
func timer1() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())
	go func() {
		time.Sleep(3 * time.Second)
		timer.Stop()
		fmt.Println("停止定时器")
	}()

	ok := timer.Stop()
	fmt.Println(ok)
	if !ok {
		// 阻塞
		x, _ := <-timer.C
		fmt.Println(x)
	}
}
