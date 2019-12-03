package string

import (
    "fmt"
    "strings"
    "unicode"
)

func main() {

    /*
    如果 a 小于 b ，返回 -1 ，反之返回 1 。不推荐使用这个函数，直接使用 == != > < >= <= 等一系列运算符更加直观。
     */
    a := "gopher"
    b := "hello world"
    fmt.Println(strings.Compare(a, b))
    fmt.Println(strings.Compare(a, a))
    fmt.Println(strings.Compare(b, a))
    fmt.Println(strings.EqualFold("GO", "go"))
    fmt.Println(strings.EqualFold("壹", "一"))

    fmt.Println("------------------------字符串包含-------------------------------------")

    fmt.Println(strings.ContainsAny("team", "i"))
    fmt.Println(strings.ContainsAny("failure", "u & i"))
    fmt.Println(strings.ContainsAny("in failure", "s g")) // 包含空格
    fmt.Println(strings.ContainsAny("foo", ""))
    fmt.Println(strings.ContainsAny("", ""))

    fmt.Println("-------------------------字符统计------------------------------------")

    /*
    字符串匹配算法：
        朴素匹配算法
        KMP 算法
        Rabin-Karp 算法
        Boyer-Moore 算法
     */
    fmt.Println(strings.Count("cheese", "e"))
    fmt.Println(len("谷歌中国"))
    fmt.Println(strings.Count("谷歌中国", ""))

    fmt.Println("-------------------------字符串分割为[]string------------------------------------")
    /*
     Fields 和 FieldsFunc、Split 和 SplitAfter、SplitN 和 SplitAfterN
    */
    // 空格分隔
    fmt.Printf("Fields are: %q \n", strings.Fields("  foo bar  baz   "))
    fmt.Println(strings.FieldsFunc("  foo bar  baz   ", unicode.IsSpace))
    // 通过 sep 进行分割，返回[]string。如果 sep 为空，相当于分成一个个的 UTF-8 字符
    // Split 会将 s 中的 sep 去掉，而 SplitAfter 会保留 sep
    fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
    fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
    // 控制切片个数 ["foo" "bar,baz"]
    fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", 2))

    fmt.Println("-------------------------字符串是否有某个前缀或后缀------------------------------------")
    /*
    HasPrefix 是否以 prefix 开始
    HasSuffix是否以 suffix 结尾
     */
    fmt.Println(strings.HasPrefix("Gopher", "Go"))
    fmt.Println(strings.HasPrefix("Gopher", "C"))
    fmt.Println(strings.HasPrefix("Gopher", ""))
    fmt.Println(strings.HasSuffix("Amigo", "go"))
    fmt.Println(strings.HasSuffix("Amigo", "Ami"))
    fmt.Println(strings.HasSuffix("Amigo", ""))

    /*
    // 在 s 中查找 sep 的第一次出现，返回第一次出现的索引
    func Index(s, sep string) int
    // 在 s 中查找字节 c 的第一次出现，返回第一次出现的索引
    func IndexByte(s string, c byte) int
    // chars 中任何一个 Unicode 代码点在 s 中首次出现的位置
    func IndexAny(s, chars string) int
    // 查找字符 c 在 s 中第一次出现的位置，其中 c 满足 f(c) 返回 true
    func IndexFunc(s string, f func(rune) bool) int
    // Unicode 代码点 r 在 s 中第一次出现的位置
    func IndexRune(s string, r rune) int

    // 有三个对应的查找最后一次出现的位置
    func LastIndex(s, sep string) int
    func LastIndexByte(s string, c byte) int
    func LastIndexAny(s, chars string) int
    func LastIndexFunc(s string, f func(rune) bool) int
     */
    fmt.Println("-------------------------字符或子串在字符串中出现的位置------------------------------------")
    han := func(c rune) bool {
        return unicode.Is(unicode.Han, c) // 汉字
    }
    fmt.Println(strings.IndexFunc("Hello, world", han))
    fmt.Println(strings.IndexFunc("Hello, 世界", han))

    fmt.Println("-------------------------字符串 JOIN 操作------------------------------------")
    fmt.Println(strings.Join([]string{"name=xxx", "age=xx"}, "&"))

    fmt.Println("-------------------------字符串重复几次------------------------------------")
    fmt.Println("ba" + strings.Repeat("na", 2))

    fmt.Println("-------------------------字符替换------------------------------------")
    mapping := func(r rune) rune {
        switch {
        case r >= 'A' && r <= 'Z': // 大写字母转小写
            return r + 32
        case r >= 'a' && r <= 'z': // 小写字母不处理
            return r
        case unicode.Is(unicode.Han, r): // 汉字换行
            return '\n'
        }
        return -1 // 过滤所有非字母、汉字的字符
    }
    fmt.Println(strings.Map(mapping, "Hello你#￥%……\n（'World\n,好Hello^(&(*界gopher..."))

    fmt.Println("-------------------------字符串子串替换------------------------------------")
    /*
    进行字符串替换时，考虑到性能问题，能不用正则尽量别用，应该用这里的函数。
    // 用 new 替换 s 中的 old，一共替换 n 个。
    // 如果 n < 0，则不限制替换次数，即全部替换
    func Replace(s, old, new string, n int) string
    // 该函数内部直接调用了函数 Replace(s, old, new , -1)
    func ReplaceAll(s, old, new string) string
     */
    fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
    fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))
    fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))

    fmt.Println("-------------------------大小写转换------------------------------------")
    // ToLower,ToUpper 用于大小写转换。ToLowerSpecial,ToUpperSpecial 可以转换特殊字符的大小写。
    fmt.Println(strings.ToLower("HELLO WORLD"))
    fmt.Println(strings.ToLower("Ā Á Ǎ À"))
    fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "壹"))
    fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "HELLO WORLD"))
    fmt.Println(strings.ToLower("Önnek İş"))
    fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))

    fmt.Println(strings.ToUpper("hello world"))
    fmt.Println(strings.ToUpper("ā á ǎ à"))
    fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "一"))
    fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "hello world"))
    fmt.Println(strings.ToUpper("örnek iş"))
    fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))

    fmt.Println("-------------------------标题处理------------------------------------")
    // 其中 Title 会将 s 每个单词的首字母大写，不处理该单词的后续字符。ToTitle 将 s 的每个字母大写。ToTitleSpecial 将 s 的每个字母大写，并且会将一些特殊字母转换为其对应的特殊大写字母。
    fmt.Println(strings.Title("hElLo wOrLd"))
    fmt.Println(strings.ToTitle("hElLo wOrLd"))
    fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "hElLo wOrLd"))
    fmt.Println(strings.Title("āáǎà ōóǒò êēéěè"))
    fmt.Println(strings.ToTitle("āáǎà ōóǒò êēéěè"))
    fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "āáǎà ōóǒò êēéěè"))
    fmt.Println(strings.Title("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
    fmt.Println(strings.ToTitle("dünyanın ilk borsa yapısı Aizonai kabul edilir"))
    fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))

    fmt.Println("-------------------------修剪------------------------------------")
    /*
    // 将 s 左侧和右侧中匹配 cutset 中的任一字符的字符去掉
    func Trim(s string, cutset string) string
    // 将 s 左侧的匹配 cutset 中的任一字符的字符去掉
    func TrimLeft(s string, cutset string) string
    // 将 s 右侧的匹配 cutset 中的任一字符的字符去掉
    func TrimRight(s string, cutset string) string
    // 如果 s 的前缀为 prefix 则返回去掉前缀后的 string , 否则 s 没有变化。
    func TrimPrefix(s, prefix string) string
    // 如果 s 的后缀为 suffix 则返回去掉后缀后的 string , 否则 s 没有变化。
    func TrimSuffix(s, suffix string) string
    // 将 s 左侧和右侧的间隔符去掉。常见间隔符包括：'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL)
    func TrimSpace(s string) string
    // 将 s 左侧和右侧的匹配 f 的字符去掉
    func TrimFunc(s string, f func(rune) bool) string
    // 将 s 左侧的匹配 f 的字符去掉
    func TrimLeftFunc(s string, f func(rune) bool) string
    // 将 s 右侧的匹配 f 的字符去掉
    func TrimRightFunc(s string, f func(rune) bool) string
     */
    x := "!!!@@@你好,!@#$ Gophers###$$$"
    fmt.Println(strings.Trim(x, "@#$!%^&*()_+=-"))
    fmt.Println(strings.TrimLeft(x, "@#$!%^&*()_+=-"))
    fmt.Println(strings.TrimRight(x, "@#$!%^&*()_+=-"))
    fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
    fmt.Println(strings.TrimPrefix(x, "!"))
    fmt.Println(strings.TrimSuffix(x, "$"))
    f := func(r rune) bool {
        return !unicode.Is(unicode.Han, r) // 非汉字返回 true
    }
    fmt.Println(strings.TrimFunc(x, f))
    fmt.Println(strings.TrimLeftFunc(x, f))
    fmt.Println(strings.TrimRightFunc(x, f))

    fmt.Println("-------------------------Reader 类型------------------------------------")
    fmt.Println("-------------------------Builder 类型------------------------------------")
    /*
    该类型实现了 io 包下的 Writer, ByteWriter, StringWriter 等接口，可以向该对象内写入数据，Builder 没有实现 Reader 等接口，所以该类型不可读，但提供了 String 方法可以获取对象内的数据。

    // 该方法向 b 写入一个字节
    func (b *Builder) WriteByte(c byte) error
    // WriteRune 方法向 b 写入一个字符
    func (b *Builder) WriteRune(r rune) (int, error)
    // WriteRune 方法向 b 写入字节数组 p
    func (b *Builder) Write(p []byte) (int, error)
    // WriteRune 方法向 b 写入字符串 s
    func (b *Builder) WriteString(s string) (int, error)
    // Len 方法返回 b 的数据长度。
    func (b *Builder) Len() int
    // Cap 方法返回 b 的 cap。
    func (b *Builder) Cap() int
    // Grow 方法将 b 的 cap 至少增加 n (可能会更多)。如果 n 为负数，会导致 panic。
    func (b *Builder) Grow(n int)
    // Reset 方法将 b 清空 b 的所有内容。
    func (b *Builder) Reset()
    // String 方法将 b 的数据以 string 类型返回。
    func (b *Builder) String() string
     */
    bu := strings.Builder{}
    _ = bu.WriteByte('7')
    n, _ := bu.WriteRune('夕')
    fmt.Println(n)
    n, _ = bu.Write([]byte("Hello, World"))
    fmt.Println(n)
    n, _ = bu.WriteString("你好，世界")
    fmt.Println(n)
    fmt.Println(bu.Len())
    fmt.Println(bu.Cap())
    bu.Grow(100)
    fmt.Println(bu.Len())
    fmt.Println(bu.Cap())
    fmt.Println(bu.String())
    bu.Reset()
    fmt.Println(bu.String())
}
