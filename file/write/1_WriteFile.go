package main
import (
	"fmt"
	"io/ioutil"
)

// 使用ioutil.WriteFile方式写入文件
// 将[]byte内容写入文件,如果content字符串中没有换行符的话，默认就不会有换行符
func main() {
	name := "./test.txt"
	content := "Hello, World!"
	data :=  []byte(content)
	if ioutil.WriteFile(name,data,0644) == nil {
		fmt.Println("写入文件成功:",content)
	}
}
