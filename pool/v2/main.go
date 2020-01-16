package main

import (
    "fmt"
    "time"
)

// -------------------------------------------
var (
    // 保存发送的工作数据
    jobs = make(chan int, 1)
)

// 工作池
type Worker struct {
}

func NewWorker() *Worker {
    return new(Worker)
}
func (w *Worker) Start(workerNum int, task func(int, int)) {
    for i := 1; i <= workerNum; i++ {
        go w.Execute(i, jobs, task)
    }
}

func (w *Worker) Execute(id int, jobs <-chan int, task func(int, int) ()) {
    for {
        task(id, <-jobs)
    }
}

// 测试
func main() {

    // task任务
    task := func(workerId, result int) {
        time.Sleep(100*time.Millisecond)
        fmt.Println("workerId", workerId, "result:", result)
    }
    NewWorker().Start(4, task)

    // 业务测试
    for j := 1; j <= 20; j++ {
        jobs <- j
    }

    for {
        time.Sleep(1)
    }
}
