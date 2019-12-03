package main

import (
	"fmt"
	"time"
)

// 定时任务
func main() {

	stopChan := make(chan bool)
	ticker := time.NewTicker(2 * time.Second)

	go func() {
		// 10秒后结束,并关闭定时器
		time.Sleep(10 * time.Second)
		stopChan <- false
	}()

	for {
		select {
		case t := <-ticker.C: // 获取时间
			fmt.Println(t)
		case ok := <-stopChan: // 判断是否需要停止定时任务
			if !ok {
				ticker.Stop()
				fmt.Println("停止任务")
				goto TICKER_END
			}
		}
	}
TICKER_END:

	fmt.Println("执行结束")
}
