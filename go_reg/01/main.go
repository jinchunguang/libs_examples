package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func regString(compile, search string) {

	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(compile)
	if reg == nil {
		fmt.Println("regexp err")
		return
	}

	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(search, -1)
	fmt.Println("regTest = ", result)
}

func regMultipleString(compile, search string) {

	//解析正则表达式，如果成功返回解释器
	reg := regexp.MustCompile(compile)
	if reg == nil {
		fmt.Println("regexp err")
		return
	}

	//根据规则提取关键信息
	result := reg.FindAllStringSubmatch(search, -1)
	//过滤<></>
	for _, text := range result {
		fmt.Println("text[1] = ", text[1])
	}
}

func regReplaceString()  {
	//目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"          //正则
	f := func(s string) string{
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v * 2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为 "##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
}
func main() {

	fmt.Println("匹配指定类型的字符串")
	regString(`a.c`, `abc azc a7c aac 888 a9c tac`)

	fmt.Println("匹配 a 和 c 中间包含一个数字的字符串")
	regString(`a[0-9]c`, `abc azc a7c aac 888 a9c tac`)

	fmt.Println("使用 \\d 来匹配 a 和 c 中间包含一个数字的字符串")
	regString(`a\dc`, `abc azc a7c aac 888 a9c tac`)

	fmt.Println("匹配字符串中的小数")
	regString(`\d+\.\d+`, `43.14 567 agsdg 1.23 7. 8.9 1sdljgl 6.66 7.8`)

	fmt.Println("匹配 div 标签中的内容")
	regMultipleString(`<div>(?s:(.*?))</div>`, `
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <title>C语言中文网 | Go语言入门教程</title>
</head>
<body>
    <div>Go语言简介</div>
    <div>Go语言基本语法
    Go语言变量的声明
    Go语言教程简明版
    </div>
    <div>Go语言容器</div>
    <div>Go语言函数</div>
</body>
</html>`)

	// 通过 Compile 方法返回一个 Regexp 对象，实现匹配，查找，替换相关的功能。
	regReplaceString()
}
