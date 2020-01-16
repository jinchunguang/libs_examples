package main

import (
	"fmt"
	"io"
	"os"
)

// WriteString()函数，用来将字符串写入一个Writer对象中
func main() {

	name := "./test.txt"
	content := "Hello, World!"

	/*
		O_RDONLY int = syscall.O_RDONLY // 只读打开文件和os.Open()同义
		O_WRONLY int = syscall.O_WRONLY // 只写打开文件
		O_RDWR   int = syscall.O_RDWR   // 读写方式打开文件
		O_APPEND int = syscall.O_APPEND // 当写的时候使用追加模式到文件末尾
		O_CREATE int = syscall.O_CREAT  // 如果文件不存在，此案创建
		O_EXCL   int = syscall.O_EXCL   // 和O_CREATE一起使用, 只有当文件不存在时才创建
		O_SYNC   int = syscall.O_SYNC   // 以同步I/O方式打开文件，直接写入硬盘.
		O_TRUNC  int = syscall.O_TRUNC  // 如果可以的话，当打开文件时先清空文件
	*/
	fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	if _, err := io.WriteString(fileObj, content); err == nil {
		fmt.Println("写入文件成功:", content)
	}
}
