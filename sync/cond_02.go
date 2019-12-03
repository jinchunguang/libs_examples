package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 生产者，消费者模型
func main()  {
	c := sync.NewCond(&sync.Mutex{})
	queueList := make(chan int, 10)

	//producer
	go func() {
		for {
			c.L.Lock()

			if len(queueList) == 2 {
				c.Wait()
			}

			num := rand.Intn(100)

			fmt.Println("producer:", num)
			queueList <- num

			if len(queueList) == 2 {
				c.Signal()
			}
			c.L.Unlock()
			time.Sleep(1*time.Second)
		}
	}()

	//consumer
	go func() {
		for {
			c.L.Lock()

			if len(queueList) == 0 {
				c.Wait()
			}

			num := <- queueList

			fmt.Println("consumer:", num)

			if len(queueList) == 0 {
				c.Signal()
			}

			c.L.Unlock()
		}
	}()

	for {
		;
	}
}