package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 使用os.Open和bufio.Reader读取文件内容
func main() {

	file := "/usr/local/nginx/logs/access.log"
	if fileObj, err := os.Open(file); err == nil {
		defer fileObj.Close()

		// 一个文件对象本身是实现了io.Reader的
		// 使用bufio.NewReader去初始化一个Reader对象，存在buffer中的，读取一次就会被清空
		reader := bufio.NewReader(fileObj)

		// 使用ReadString(delim byte)来读取delim以及之前的数据并返回相关的字符串.
		if result, err := reader.ReadString(byte('\n')); err == nil {
			fmt.Println("使用ReadSlince相关方法读取内容:", result)
		}

		// 注意:上述ReadString已经将buffer中的数据读取出来了，下面将不会输出内容
		// 需要注意的是，因为是将文件内容读取到[]byte中，因此需要对大小进行一定的把控
		buf := make([]byte, 1024)

		// 读取Reader对象中的内容到[]byte类型的buf中
		if n, err := reader.Read(buf); err == nil {
			fmt.Println(strconv.Itoa(n))
			fmt.Println(string(buf))
		}

	}
}
