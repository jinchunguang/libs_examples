package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var err error
var file *os.File

func main() {

	/*
	   打开和关闭文件
	*/
	fmt.Println("---------------------------打开关闭-----------------------------------")
	// 只读方式打开当前目录下的main.go文件
	file, err = os.Open("./Write.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	// 关闭文件
	file.Close()

	/*
	   读取文件
	   它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回0和io.EOF。
	*/
	fmt.Println("---------------------------读取文件-----------------------------------")
	// 只读方式打开当前目录下的main.go文件
	file, err := os.Open("./Write.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 使用Read方法读取数据
	var tmp = make([]byte, 2048)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		return
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节数据\n", n)
	fmt.Println(string(tmp[:n]))

	/*
	   循环读取
	*/
	fmt.Println("---------------------------循环读取-----------------------------------")
	// 只读方式打开当前目录下的main.go文件
	file, err = os.Open("./Write.txt")
	if err != nil {
		fmt.Println("open file failed!, err:", err)
		return
	}
	defer file.Close()
	// 循环读取文件
	var content []byte
	var tmp1 = make([]byte, 2048)
	for {
		n, err := file.Read(tmp1)
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))

	/*
	   bufio读取文件
	   bufio 是在file的基础上封装了一层API，支持更多的功能。
	*/
	fmt.Println("---------------------------bufio读取-----------------------------------")
	file, err = os.Open("./Write.txt")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		fmt.Print(line)
	}

	/*
	   ioutil读取整个文件
	*/
	fmt.Println("---------------------------bufio读取-----------------------------------")
	content, err = ioutil.ReadFile("./Write.txt")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))

	/*
	   文件写入操作
	*/
	fmt.Println("---------------------------Write和WriteString-----------------------------------")
	file, err = os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "hello 沙河 "
	file.Write([]byte(str))       //写入字节切片数据
	file.WriteString("hello 小王子") //直接写入字符串数据

	fmt.Println("---------------------------bufio.NewWriter-----------------------------------")
	file, err = os.OpenFile("test.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("hello,hello,hello\n") //将数据先写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件

	fmt.Println("---------------------------ioutil.WriteFile-----------------------------------")
	str = "hello 沙河"
	err = ioutil.WriteFile("./test.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
