/**
 * @Author : jinchunguang
 * @Date : 19-10-31 上午10:26
 * @Project : sty
 */
package time

import (
	"fmt"
	"time"
)

func main() {
	/*// 新建计时器，120秒以后触发，go触发计时器的方法比较特别，就是在计时器的channel中发送值
	  tick := time.NewTicker(3 * time.Second)

	  for {
	      select {
	      // 此处在等待channel中的信号，因此执行此段代码时会阻塞120秒
	      case <-tick.C:
	          log.Println(time.Now())
	      default:
	          time.Sleep(100*time.Millisecond)
	      }

	  }*/

	start := time.Now()
	timer := time.AfterFunc(2*time.Second, func() {
		fmt.Println("after func callback, elaspe:", time.Now().Sub(start))
	})

	time.Sleep(1 * time.Second)
	// time.Sleep(3*time.Second)
	// Reset 在 Timer 还未触发时返回 true；触发了或Stop了，返回false
	if timer.Reset(3 * time.Second) {
		fmt.Println("timer has not trigger!")
	} else {
		fmt.Println("timer had expired or stop!")
	}

	time.Sleep(10 * time.Second)

	// output:
	// timer has not trigger!
	// after func callback, elaspe: 4.00026461s
}
