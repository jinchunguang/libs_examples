package main

import (
	"fmt"
	"sync"
	"time"
)

/**
RWMutex是读写互斥锁。

该锁可以被同时多个读取者持有或唯一个写入者持有。
RWMutex可以创建为其他结构体的字段；零值为解锁状态。
RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。

	// 读写互斥
	var rwx sync.RWMutex

	// 写模式
	rwx.Lock() // 加锁，一路只写
	rwx.Unlock() // 释放锁

	// 读模式
	rwx.RLock() // 加锁，多路只读
	rwx.RUnlock()// 释放锁
 */

func main() {

	var wg sync.WaitGroup
	var rwx sync.RWMutex
	for i:=0;i<5;i++ {

		// 读数据,多路只读
		wg.Add(1)
		go func(i int) {

			rwx.RLock()

			fmt.Println("读数据库...",i)
			<-time.After(2*time.Second)

			rwx.RUnlock()
			defer wg.Done()
		}(i)

		// 写数据,一路只写
		wg.Add(1)
		go func(i int) {

			rwx.Lock()

			fmt.Println("写数据库...",i)
			<-time.After(2*time.Second)

			rwx.Unlock()
			defer wg.Done()
		}(i)

	}

	wg.Wait()

	fmt.Println("程序结束")

}
