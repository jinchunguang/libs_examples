package main

import (
    "fmt"

    "github.com/go-ego/gse"
)

var seg gse.Segmenter

func cut() {
    text := "你好世界, Hello world."

    hmm := seg.Cut(text, true)
    fmt.Println("cut use hmm: ", hmm)

    hmm = seg.CutSearch(text, true)
    fmt.Println("cut search use hmm: ", hmm)

    hmm = seg.CutAll(text)
    fmt.Println("cut all: ", hmm)
}

func segCut() {
    // 分词文本
    tb := []byte("山达尔星联邦共和国联邦政府")

    // 处理分词结果
    // 支持普通模式和搜索模式两种分词，见代码中 ToString 函数的注释。
    // 搜索模式主要用于给搜索引擎提供尽可能多的关键字
    fmt.Println("输出分词结果, 类型为字符串, 使用搜索模式: ", seg.String(tb, true))
    fmt.Println("输出分词结果, 类型为 slice: ", seg.Slice(tb))

    segments := seg.Segment(tb)
    // 处理分词结果
    fmt.Println(gse.ToString(segments))

    text1 := []byte("上海地标建筑, 东方明珠电视台塔上海中心大厦")
    segments1 := seg.Segment([]byte(text1))
    fmt.Println(gse.ToString(segments1, true))
}

func main() {
    // 加载默认字典
    //seg.LoadDict()
    // 载入词典
    //seg.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")
    seg.LoadDict("data/dict/dictionary.txt")

    cut()

    segCut()
}