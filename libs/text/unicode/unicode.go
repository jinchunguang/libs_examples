package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {

	/*
	   UTF-8 表示最少用一个字节就能表示一个字符的编码实现。它采取的方式是对不同的语言使用不同的方法，将 unicode 编码按照这个方法进行转换。 我们只要记住最后的结果是英文占一个字节，中文占三个字节。

	   unicode
	   unicode/utf8
	   unicode/utf16


	   func IsControl(r rune) bool  // 是否控制字符
	   func IsDigit(r rune) bool  // 是否阿拉伯数字字符，即 0-9
	   func IsGraphic(r rune) bool // 是否图形字符
	   func IsLetter(r rune) bool // 是否字母
	   func IsLower(r rune) bool // 是否小写字符
	   func IsMark(r rune) bool // 是否符号字符
	   func IsNumber(r rune) bool // 是否数字字符，比如罗马数字Ⅷ也是数字字符
	   func IsOneOf(ranges []*RangeTable, r rune) bool // 是否是 RangeTable 中的一个
	   func IsPrint(r rune) bool // 是否可打印字符
	   func IsPunct(r rune) bool // 是否标点符号
	   func IsSpace(r rune) bool // 是否空格
	   func IsSymbol(r rune) bool // 是否符号字符
	   func IsTitle(r rune) bool // 是否 title case
	   func IsUpper(r rune) bool // 是否大写字符
	   func Is(rangeTab *RangeTable, r rune) bool // r 是否为 rangeTab 类型的字符
	   func In(r rune, ranges ...*RangeTable) bool  // r 是否为 ranges 中任意一个类型的字符
	*/

	fmt.Println("-----------------是否控制字符------------------")
	single := '\u0015'
	fmt.Println(unicode.IsControl(single))
	single = '\ufe35'
	fmt.Println(unicode.IsControl(single))

	fmt.Println("-----------------数字字符------------------")
	digit := '1'
	fmt.Println(`是否阿拉伯数字字符:`, unicode.IsDigit(digit))
	fmt.Println(`是否数字字符，比如罗马数字Ⅷ也是数字字符:`, unicode.IsNumber(digit))

	letter := 'Ⅷ'
	fmt.Println(unicode.IsDigit(letter))
	fmt.Println(unicode.IsNumber(letter))

	han := '你'
	fmt.Println(unicode.IsDigit(han))
	fmt.Println(unicode.Is(unicode.Han, han))
	fmt.Println(unicode.In(han, unicode.Gujarati, unicode.White_Space))

	// utf8 包
	/*
	   判断是否符合 utf8 编码的函数：

	   func Valid(p []byte) bool
	   func ValidRune(r rune) bool
	   func ValidString(s string) bool
	   判断 rune 所占字节数：

	   func RuneLen(r rune) int
	   判断字节串或者字符串的 rune 数：

	   func RuneCount(p []byte) int
	   func RuneCountInString(s string) (n int)
	   编码和解码到 rune：

	   func EncodeRune(p []byte, r rune) int
	   func DecodeRune(p []byte) (r rune, size int)
	   func DecodeRuneInString(s string) (r rune, size int)
	   func DecodeLastRune(p []byte) (r rune, size int)
	   func DecodeLastRuneInString(s string) (r rune, size int)
	   是否为完整 rune：

	   func FullRune(p []byte) bool
	   func FullRuneInString(s string) bool
	   是否为 rune 第一个字节：

	   func RuneStart(b byte) bool
	*/

	fmt.Println("-----------------utf8 包------------------")

	word := []byte("界")

	fmt.Println("-----------------判断是否符合 utf8 编码的函数------------------")
	fmt.Println(utf8.Valid(word[:2]))
	fmt.Println(utf8.ValidRune('界'))
	fmt.Println(utf8.ValidString("世界"))

	fmt.Println("-----------------判断 rune 所占字节数------------------")
	fmt.Println(utf8.RuneLen('界'))

	fmt.Println("-----------------判断字节串或者字符串的 rune 数------------------")
	fmt.Println(utf8.RuneCount(word))
	fmt.Println(utf8.RuneCountInString("世界"))

	fmt.Println("-----------------编码和解码到 rune------------------")
	p := make([]byte, 3)
	utf8.EncodeRune(p, '好')
	fmt.Println(p)
	fmt.Println(utf8.DecodeRune(p))
	fmt.Println(utf8.DecodeRuneInString("你好"))
	fmt.Println(utf8.DecodeLastRune([]byte("你好")))
	fmt.Println(utf8.DecodeLastRuneInString("你好"))

	fmt.Println("-----------------是否为完整 rune------------------")
	fmt.Println(utf8.FullRune(word[:2]))
	fmt.Println(utf8.FullRuneInString("你好"))

	fmt.Println("-----------------是否为 rune 第一个字节------------------")
	fmt.Println(utf8.RuneStart(word[1]))
	fmt.Println(utf8.RuneStart(word[0]))

}
