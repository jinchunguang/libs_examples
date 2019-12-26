package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)
// ioutil.ReadAll()读取文件示例
// ReadAll从源中读取数据直到EOF
func main() {
	file := "/usr/local/nginx/logs/access.log"
	if fileObj,err := os.Open(file);err == nil {
		//if fileObj,err := os.OpenFile(name,os.O_RDONLY,0644); err == nil {
		defer fileObj.Close()
		if contents,err := ioutil.ReadAll(fileObj); err == nil {
			result := strings.Replace(string(contents),"\n","",1)
			fmt.Println("Use os.Open family functions and ioutil.ReadAll to read a file contents:",result)
		}

	}
}
