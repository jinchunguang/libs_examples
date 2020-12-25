/**
 * @Author : jinchunguang
 * @Date : 19-12-26 上午10:49
 * @Project : libs_examples
 */
package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	logFile, err := os.OpenFile("./biz.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetPrefix("[biz] ")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}

func main() {
	log.Println("这是一条文件日志。")
	log.Printf("LocalInfo %s", "这是一条文件日志。")
}
