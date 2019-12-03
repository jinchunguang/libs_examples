package main

import (
	"fmt"
	"time"
)

/**
sync包提供了基本的同步基元，如互斥锁。除了Once和WaitGroup类型，大部分都是适用于低水平程序线程，高水平的同步使用channel通信更好一些。
 */

func do3() {
	time.Sleep(3 * time.Second)
	fmt.Println("do ok")
}

// 等待组
func main() {
	go do3()
	fmt.Println("main ok")

	// 直接输出main ok
}
