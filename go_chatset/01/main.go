/*
Golang 中的 UTF-8 与 GBK 编码转换
 */

package main

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
)

// GBK 转 UTF-8
func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// UTF-8 转 GBK
func Utf8ToGbk(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

// 编码转换测试
func main() {

	s := "[debug]: 编码转换测试"

	gbk, err := Utf8ToGbk([]byte(s))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(gbk))
	}

	utf8, err := GbkToUtf8(gbk)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(utf8))
	}
}