package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// 使用ReadFile直接读取
// 读取文件内容，并返回[]byte数据和错误信息。err == nil时，读取成功
func main() {
	file := "/usr/local/nginx/logs/access.log"
	if contents, err := ioutil.ReadFile(file); err == nil {
		//[]byte类型，转换成string类型后会多一行空格,使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		fmt.Println(result)
	}
}
