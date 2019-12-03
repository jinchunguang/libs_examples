package bak

import (
	"fmt"
	"io/ioutil"
	"os"
)
func main() {
	//IoutilReadFile("test.txt")
	OsReadAll("test.txt")
}

// ioutil.ReadFile
func IoutilReadFile(name string) {
	contents,err := ioutil.ReadFile(name);
	if err!=nil{
		fmt.Println(err)
	}
	// []byte类型转换成string
	result := string(contents)
	fmt.Println(result)
}

// ioutil.ReadAll
func OsReadAll(name string) {
	file,err := os.Open(name);
	if err!=nil{
		fmt.Println(err)
	}
	result,err:=ioutil.ReadAll(file)
	// []byte类型转换成string
	fmt.Println(string(result))
}
