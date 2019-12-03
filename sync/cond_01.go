package main

import (
	"fmt"
	"sync"
	"time"
)

var cond = sync.NewCond(&sync.Mutex{})
func main() {

	houseRisePrice:=true

	go func() {
		for   {
			time.Sleep(3*time.Second)
			cond.L.Lock()

			// 给监听者发送信号
			// cond.Broadcast 发送通知，唤醒所有的线程
			// cond.Signal 发送广播，唤醒一个线程（如果存在）
			houseRisePrice=false
			cond.Signal()
			cond.L.Unlock()
		}
	}()

	for   {
		go func() {
			cond.L.Lock()
			if houseRisePrice {
				fmt.Println("可以观望")
				// 解锁并阻塞当前线程
				// Wait除非被Broadcast或者Signal唤醒，不会主动返回。
				cond.Wait()
			}

			fmt.Println("谨慎买房")
			cond.L.Unlock()
		}()
		time.Sleep(1*time.Second)
	}


	fmt.Println("main over!")

}
