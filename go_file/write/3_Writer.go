package main

import (
	"bufio"
	"fmt"
	"os"
)

// 使用bufio包中Writer对象的相关方法进行数据的写入
func main() {

	name := "./test.txt"
	content := "Hello, World!\n"

	if fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644);err == nil {
		defer fileObj.Close()

		// 使用WriteString方法,写入字符串并返回写入字节数和错误信息
		writeObj := bufio.NewWriterSize(fileObj,4096)
		if _,err := writeObj.WriteString(content);err == nil {
			fmt.Println(content)
		}

		// 使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
		buf := []byte(content)
		if _,err := writeObj.Write(buf);err == nil {
			if  err := writeObj.Flush(); err != nil {panic(err)}
			fmt.Println(content)
		}
	}
}
