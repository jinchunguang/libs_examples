/*
 go 工作池
*/

package v1

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, f func(int) ()) {
    for j := range jobs {
        fmt.Println("workerId:", id, "businessID:", j)
        // 执行任务
        f(j * 10)
    }
}

func main() {

    f := func(result int) {
        fmt.Println("result:", result)
    }
    /*
     为了使用我们的工作池 我们需要发送工作和接收结果的
    */
    jobs := make(chan int, 1)

    for w := 1; w <= 100; w++ {
        go worker(w, jobs, f)
    }

    // 业务测试
    for j := 1; j <= 100; j++ {
        jobs <- j
    }

    close(jobs)

    for  {
        time.Sleep(1)
    }
}
