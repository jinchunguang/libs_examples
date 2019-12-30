package main

import (
    "fmt"
    "html/template"
    "io/ioutil"
    "net/http"
)

// main.go

type UserInfo struct {
    Name   string
    Gender string
    Age    int
}

func sayHello3(w http.ResponseWriter, r *http.Request) {
    htmlByte, err := ioutil.ReadFile("./hello3.tmpl")
    if err != nil {
        fmt.Println("read html failed, err:", err)
        return
    }
    // 自定义一个夸人的模板函数
    kua := func(arg string) (string, error) {
        return arg + "真帅", nil
    }
    // 采用链式操作在Parse之前调用Funcs添加自定义的kua函数
    tmpl, err := template.New("hello").Funcs(template.FuncMap{"kua": kua}).Parse(string(htmlByte))
    if err != nil {
        fmt.Println("create template failed, err:", err)
        return
    }

    user := UserInfo{
        Name:   "小王子",
        Gender: "男",
        Age:    18,
    }
    // 使用user渲染模板，并将结果写入w
    tmpl.Execute(w, user)
}

func sayHello2(w http.ResponseWriter, r *http.Request) {
    // 解析指定文件生成模板对象
    tmpl, err := template.ParseFiles("./hello2.tmpl")
    if err != nil {
        fmt.Println("create template failed, err:", err)
        return
    }
    // 利用给定数据渲染模板，并将结果写入w
    user := UserInfo{
        Name:   "小王子",
        Gender: "男",
        Age:    18,
    }
    tmpl.Execute(w, user)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
    // 解析指定文件生成模板对象
    tmpl, err := template.ParseFiles("./hello.tmpl")
    if err != nil {
        fmt.Println("create template failed, err:", err)
        return
    }
    // 利用给定数据渲染模板，并将结果写入w
    tmpl.Execute(w, "小王子")
}

func tmplDemo(w http.ResponseWriter, r *http.Request) {
    tmpl, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
    if err != nil {
        fmt.Println("create template failed, err:", err)
        return
    }
    user := UserInfo{
        Name:   "小王子",
        Gender: "男",
        Age:    18,
    }
    tmpl.Execute(w, user)
}


func main() {
    http.HandleFunc("/sh", sayHello)
    http.HandleFunc("/sh2", sayHello2)
    http.HandleFunc("/sh3", sayHello3)
    http.HandleFunc("/tmpl", tmplDemo)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        fmt.Println("HTTP server failed,err:", err)
        return
    }
}