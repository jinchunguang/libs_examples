package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("------------------------------------------")
	fmt.Println(time.Local.String())
	// 得到了 sec 和 nsec
	fmt.Println(time.Now())

	fmt.Println("----------------与 Unix 时间戳的转换--------------------------")
	/*
	   time.Unix(sec, nsec int64) 通过 Unix 时间戳生成 time.Time 实例；
	   time.Time.Unix() 得到 Unix 时间戳；
	   time.Time.UnixNano() 得到 Unix 时间戳的纳秒表示
	*/
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())

	fmt.Println("----------------格式化和解析--------------------------")
	/*
	   time.Parse 和 time.ParseInLocation
	   time.Time.Format
	*/

	t, _ := time.Parse("2006-01-02 15:04:05", "2016-06-13 09:14:00")
	fmt.Println(time.Now().Sub(t).Hours())

	t, _ = time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 09:14:00", time.Local)
	fmt.Println(time.Now().Sub(t).Hours())

	fmt.Println("----------------Round 和 Truncate 方法--------------------------")

	// t, _ = time.ParseInLocation("2006-01-02 15:04:05", "2016-06-13 15:34:39", time.Local)
	t, _ = time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	// 整点（向下取整）
	fmt.Println(t.Truncate(1 * time.Hour))
	// 整点（最接近）
	fmt.Println(t.Round(1 * time.Hour))

	// 整分（向下取整）
	fmt.Println(t.Truncate(1 * time.Minute))
	// 整分（最接近）
	fmt.Println(t.Round(1 * time.Minute))

	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
	fmt.Println(t2)

	fmt.Println("------------------------通过 time.After 模拟超时----------------------------------")
	c := make(chan int)

	go func() {
		// time.Sleep(1 * time.Second)
		time.Sleep(3 * time.Second)
		<-c
	}()

	select {
	case c <- 1:
		fmt.Println("channel...")
	case <-time.After(2 * time.Second):
		close(c)
		fmt.Println("timeout...")
	}

	fmt.Println("------------------------time.Stop 停止定时器 或 time.Reset 重置定时器----------------------------------")
	start := time.Now()
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after func callback, elaspe:", time.Now().Sub(start))
	})

	time.Sleep(1 * time.Second)
	// time.Sleep(3*time.Second)
	// Reset 在 Timer 还未触发时返回 true；触发了或 Stop 了，返回 false
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	}

	fmt.Println("------------------------time.NewTimer----------------------------------")
	t1 := time.NewTimer(time.Second * 2)

	for {
		select {
		case <-t1.C:
			println("5s timer")
		default:

		}
	}

}
