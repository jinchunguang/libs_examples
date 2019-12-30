package main

import (
    "bytes"
    "compress/flate"
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {

    str:="test123"
    b:=Gzdeflate(str,-1)
    ss:=Gzdecode(string(b))
    fmt.Println(ss)
}

// 解码
func Gzdecode(data string) string  {
    if data == "" {
        return ""
    }
    r :=flate.NewReader(strings.NewReader(data))
    defer r.Close()
    out, err := ioutil.ReadAll(r)
    if err !=nil {
        fmt.Errorf("%s\n",err)
        return ""
    }
    return string(out)
}

// 编码
func Gzdeflate(data string,level int) []byte  {
    if data == "" {
        return []byte{}
    }
    var bufs bytes.Buffer
    w,_ :=flate.NewWriter(&bufs,level)
    w.Write([]byte(data))
    w.Flush()
    defer w.Close()
    return bufs.Bytes()
}

// 编码
func GzdeflateForString(data string,level int) string  {
    if data == "" {
        return ""
    }
    var bufs bytes.Buffer
    w,_ :=flate.NewWriter(&bufs,level)
    w.Write([]byte(data))
    w.Flush()
    defer w.Close()
    return bufs.String()
}