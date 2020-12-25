package Split

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	/*
	   分词
	*/
	const input = "This is The Golang Standard Library.\nWelcome you!"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		count++
		// 遍历每一个分词
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	// 词汇总数
	fmt.Println(count)

	/*
	   将文本分词并且写入到文件
	*/
	file, err := os.Create("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")
	// 将文件 offset 设置到文件开头
	// file.Seek(0, os.SEEK_SET) // 丢弃
	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
