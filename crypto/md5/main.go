package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
)

/*
字符串md5
 */
func stringMd5(str string) string {
	m := md5.New()
	_, err := io.WriteString(m, str)
	if err != nil {
		log.Fatal(err)
	}
	arr := m.Sum(nil)
	return fmt.Sprintf("%x", arr)
}

/*
字节md5
 */
func byteMd5(str []byte) string {
	return fmt.Sprintf("%x", md5.Sum(str))
}

func main() {
	str := "111111"
	fmt.Println(stringMd5(str))
	fmt.Println(byteMd5([]byte(str)))
}
