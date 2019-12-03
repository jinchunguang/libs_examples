package main

import (
	"fmt"
	"sync"
)

/*
sync.Mutex 互斥锁
	mx.Lock() 单协程加锁，其它协程只能等待该锁释放。抢到锁，执行；没抢到锁，阻塞。资源竞争
	mx.Unlock() 单协程释放锁
*/
// 线程安全加互斥锁
func main() {
	var wg sync.WaitGroup

	var mx sync.Mutex

	money := 2000

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func(i int) {

			mx.Lock()

			for j := 1; j <= 10; j++ {
				money++
			}

			mx.Unlock()
			defer wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(money)
}
