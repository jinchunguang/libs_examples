/**
 * @Author : jinchunguang
 * @Date : 19-12-26 上午10:37
 * @Project : libs_examples
 */
package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    /*
    log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。
    func Flags() int
    func SetFlags(flag int)

    flag选项
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
     */
    log.SetFlags(log.Lshortfile | log.LstdFlags)
    log.Println("这是一条很普通的日志。")

    /*
    log标准库中还提供了关于日志信息前缀的两个方法：
    func Prefix() string
    func SetPrefix(prefix string)
     */
    log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
    log.SetPrefix("[SRV] ")
    log.Println("这是一条很普通的日志。")


    /*
    配置日志输出位置
    func SetOutput(w io.Writer)
    SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。
     */
    logFile, err := os.OpenFile("./biz.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("open log file failed, err:", err)
        return
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Lshortfile | log.Lmicroseconds | log.Ldate)
    log.SetPrefix("[biz] ")
    log.Println("这是一条文件日志。")

}
