package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("-----------------字符串转为整型------------------")

	/*

	   func ParseInt(s string, base int, bitSize int) (i int64, err error)
	   func ParseUint(s string, base int, bitSize int) (n uint64, err error)
	   func Atoi(s string) (i int, err error)

	   ParseInt、ParseUint 和 Atoi
	   Atoi 是 ParseInt 的便捷版，内部通过调用 ParseInt(s, 10, 0) 来实现的；
	   ParseInt 转为有符号整型；
	   ParseUint 转为无符号整型。

	   参数 base 代表字符串按照给定的进制进行解释。
	   一般的，base 的取值为 2~36，
	   如果 base 的值为 0，则会根据字符串的前缀来确定 base 的值："0x" 表示 16 进制； "0" 表示 8 进制；否则就是 10 进制。

	   参数 bitSize 表示的是整数取值范围，或者说整数的具体类型。取值 0、8、16、32 和 64 分别代表 int、int8、int16、int32 和 int64。

	   Go 中，int/uint 类型，不同系统能表示的范围是不一样的，
	   32 位系统占 4 个字节；64 位系统占 8 个字节。
	   当 bitSize==0 时，应该表示 32 位还是 64 位呢？这里没有利用 runtime.GOARCH 之类的方式，而是巧妙的通过如下表达式确定 intSize：
	*/

	// n, err := strconv.ParseInt("128", 10, 8)
	n, err := strconv.ParseInt("128", 10, 0)
	fmt.Println(n, err)
	m, err := strconv.Atoi("128")
	fmt.Println(m, err)

	fmt.Println("-----------------整型转为字符串------------------")
	/*
	   func FormatUint(i uint64, base int) string    // 无符号整型转字符串
	   func FormatInt(i int64, base int) string    // 有符号整型转字符串
	   func Itoa(i int) string

	   FormatInt: 性能好一些
	*/
	x := strconv.Itoa(128)
	fmt.Println(x)
	x = strconv.FormatInt(int64(128), 10)
	fmt.Println(x)

	// startTime := time.Now()
	// for i := 0; i < 10000; i++ {
	//     fmt.Sprintf("%d", i)
	// }
	// fmt.Println(time.Now().Sub(startTime))
	//
	// startTime = time.Now()
	// for i := 0; i < 10000; i++ {
	//     strconv.Itoa(i)
	// }
	// fmt.Println(time.Now().Sub(startTime))

	fmt.Println("-----------------字符串和布尔值之间的转换------------------")
	// // 接受 1, t, T, TRUE, true, True, 0, f, F, FALSE, false, False 等字符串；
	// // 其他形式的字符串会返回错误
	// func ParseBool(str string) (value bool, err error)
	// // 直接返回 "true" 或 "false"
	// func FormatBool(b bool) string
	// // 将 "true" 或 "false" append 到 dst 中
	// // 这里用了一个 append 函数对于字符串的特殊形式：append(dst, "true"...)
	// func AppendBool(dst []byte, b bool)

	b, err := strconv.ParseBool("0")
	fmt.Println(b, err)
	fmt.Println(strconv.FormatBool(true))

	fmt.Println("-----------------字符串和浮点数之间的转换------------------")
	//
	// func ParseFloat(s string, bitSize int) (f float64, err error)
	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// func AppendFloat(dst []byte, f float64, fmt byte, prec int, bitSize int)
	/*
	   而 prec 表示有效数字（对 fmt='b' 无效），对于 'e', 'E' 和 'f'，有效数字用于小数点之后的位数；
	   对于 'g' 和 'G'，则是所有的有效数字。例如：
	*/

	fmt.Println(strconv.FormatFloat(1223.13252, 'e', 3, 32)) // 结果：1.223e+03
	fmt.Println(strconv.FormatFloat(1223.13252, 'g', 3, 32)) // 结果：1.22e+03

	s := strconv.FormatFloat(1234.5678, 'g', 6, 64)
	f6, err := strconv.ParseFloat(s, 64)
	fmt.Println(f6, err)

	fmt.Println("-----------------其他导出的函数------------------")
	fmt.Println(`This is "studygolang.com" website`)
	fmt.Println("This is \"studygolang.com\" website")
	fmt.Println("This is", strconv.Quote("studygolang.com"), "website")
}
