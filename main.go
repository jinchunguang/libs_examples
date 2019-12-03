/**
 * @Author : jinchunguang
 * @Date : 19-11-2 下午7:00
 * @Project : sty
 */
package main

import (
    "fmt"
    "github.com/jinchunguang/demo/adlog"
    "log"
    "os"
    "time"
)

func main() {
    l := adlog.LogStruct{}
    if err := l.InitLog(); err != nil {
        log.Println("********",err)
        os.Exit(1)
    }

    r := l.INFO(time.Now().Format("2006-01-02 15:04:05\n"))
    fmt.Println("rrrrrrrrrrrrrr: ", r)
}

