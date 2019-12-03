package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "regexp"
    "runtime"
    "sync"
    "time"
)

type Archives struct {
    Title string
    Time string
    content string
}


var count = 0

func fetch(urls chan <- string ,wg *sync.WaitGroup,url string) { // 获取主要数据页

    res,_ := http.Get(`http://www.ityouknow.com/archives.html`)
    defer res.Body.Close()
    d,_ := ioutil.ReadAll(res.Body)
    text := string(d)
    reg := regexp.MustCompile(`<li><span>(.*)?<a(.*)?href="(.*)">(.*)?></li>`)
    re := reg.FindAllStringSubmatch(text,-1)

    for _,v := range re {
        urls <- "http://www.ityouknow.com" + v[3]
    }

    close(urls)
}


func fetchInfo(urls <- chan string,wg *sync.WaitGroup) {
    for v := range urls {

        go func(v string,wg *sync.WaitGroup) {

            wg.Add(1)
            defer func() {
                if err := recover();err != nil {
                    fmt.Println(err,)
                }
            }()

            res,err := http.Get(v)
            if err != nil {
                return
            }
            defer res.Body.Close()

            d,_ := ioutil.ReadAll(res.Body)
            html := string(d)
            titleReg := regexp.MustCompile(`<section class="jumbotron geopattern" data-pattern-id="(.*)?"`)
            timeReg := regexp.MustCompile(`<span class="octicon octicon-calendar"></span>(.*)?|\n\s</span>`)

            tits := titleReg.FindAllStringSubmatch(html,-1)
            times := timeReg.FindAllStringSubmatch(html,-1)

            for i,v := range tits {
                if len(times) == 0 {
                    continue
                }

                fmt.Println(v[1],times[i][1])
            }
        }(v,wg)

    }

    fmt.Println("xx")
}


func main() {
    runtime.GOMAXPROCS(4)
    defer func() func(){
        start := time.Now()
        return func() {
            fmt.Println(time.Since(start))
        }
    }()()


    var urls = make(chan string,20)
    var wg = new(sync.WaitGroup)

    go fetch(urls,wg,`http://www.ityouknow.com/archives.html`)


    fetchInfo(urls,wg)


    wg.Wait()

}