package main

import (
	"fmt"
	"sync"
	"time"
)

/**
sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。

sync.WaitGroup
	wg.Add() 等待组添加协程
	wg.Wait()  等待组协程归零
	wg.Done() 等待组减少协程

*/
func do1(wg *sync.WaitGroup) {
	time.Sleep(3 * time.Second)
	fmt.Println("do ok")
	// 等待组减少协程
	defer wg.Done()
}

// 等待组
func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// 等待组添加协程
		wg.Add(1)
		go do1(&wg)
	}
	// 等待组协程归零
	wg.Wait()
	fmt.Println("main ok")
}
