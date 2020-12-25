package main

import (
	"fmt"
	"sync"
)

// 线程不安全
func main() {
	var wg sync.WaitGroup
	money := 2000

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			// 数值变大，就会出现线程不安全
			for j := 1; j <= 10; j++ {
				money++
			}
			defer wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(money)
}
