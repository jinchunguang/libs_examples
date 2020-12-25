package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 使用文件对象的Read方法读取
// Read方法从文件中读取最多len(b)字节数据并写入
func main() {
	file := "/usr/local/nginx/logs/access.log"
	if fileObj, err := os.Open(file); err == nil {
		defer fileObj.Close()
		//在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
		buf := make([]byte, 1024)
		if n, err := fileObj.Read(buf); err == nil {
			fmt.Println("字节数:" + strconv.Itoa(n))
			result := strings.Replace(string(buf), "\n", "", 1)
			fmt.Println(result)
		}
	}
}
