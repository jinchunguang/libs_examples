package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	var oc sync.Once

	count  := 20
	wg.Add(count)

	// 20个协程执行同一个任务，但是只执行一次
	fmt.Println("任务开始")
	for i := 1; i <= count; i++ {

		fmt.Println("杀手",i,"出发")
		go func(i int) {
			// 人只能死一次
			oc.Do(func() {
				fmt.Println("杀死大汉奸","杀手编号:",i)
			})
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("任务结束")
}
