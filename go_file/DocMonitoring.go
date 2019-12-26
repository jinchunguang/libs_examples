package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

// 文件监控
type DocMonitor struct {
	LogPath   string      // 日志文件路径
	ReadChan  chan string // 读取chan
	WriteChan chan string // 写入chan
}

// 文件读取逻辑
func (dm *DocMonitor) Read() {
	fmt.Println("Read start-up success")

	fileObj, err := os.Open(dm.LogPath);
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	defer fileObj.Close()

	buf := bufio.NewReader(fileObj)
	// 定位到末尾
	fileObj.Seek(0, 2)
	for {
		if result, err := buf.ReadString('\n'); err == nil {
			if err == io.EOF {
				time.Sleep(100 * time.Millisecond)
				continue;
			} else {
				dm.ReadChan <- result
			}
		}
	}

}

// 内容处理逻辑
func (dm *DocMonitor) Handle() {
	fmt.Println("Handle start-up success")
	for   {
		content := <-dm.ReadChan
		// 内容处理逻辑
		dm.WriteChan <- strings.ToLower(content)
	}
}

// 文件存储逻辑 (文件/db/nosql)
func (dm *DocMonitor) Write() {

	fmt.Println("Write start-up success")

	// 将处理后的结果写入文件
	name := "./Write.txt"
	fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}

	for   {
		content := <-dm.WriteChan
		if  _,err := io.WriteString(fileObj,content);err == nil {
			fmt.Println("写入成功:",content)
		}
	}
}

// 文件实时监控
// go run readlog.go /usr/local/nginx/logs/access.log
func main() {

	args:=os.Args

	//file := "/usr/local/nginx/logs/access.log"
	file := args[1]

	dm := &DocMonitor{
		LogPath:  file,
		ReadChan: make(chan string, 1024),
		WriteChan: make(chan string, 1024),
	}

	go dm.Read()
	go dm.Handle()
	go dm.Write()


	for {
		time.Sleep(time.Second)
	}

}


