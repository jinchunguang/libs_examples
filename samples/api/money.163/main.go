package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()
	end := now.Format("20060102")
	start := now.AddDate(0, 0, -120).Format("20060102")
	url := "http://quotes.money.163.com/service/chddata.html?code=0600519&start=" + start + "&end=" + end
	req, _ := http.NewRequest("GET", url, nil)
	// 设置header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1;WOW64) AppleWebKit/537.36 (KHTML,like GeCKO) Chrome/45.0.2454.85 Safari/537.36 115Broswer/6.0.3")
	req.Header.Set("Referer", "https://movie.douban.com/")
	req.Header.Set("Connection", "keep-alive")
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}
	csv1, _ := ioutil.ReadAll(resp.Body)
	r2 := csv.NewReader(strings.NewReader(string(csv1)))
	ss, _ := r2.ReadAll()

	sz := len(ss)
	buf := new(bytes.Buffer)
	maxv := 0.0
	minv := 4096.0
	todayp, _ := strconv.ParseFloat(ss[1][3], 64)
	for i := 1; i < sz; i++ { //第0行是标题，从第1行取数据
		buf.WriteString(" " + ss[i][3] + " ")
		buf.WriteString("\r\n")
		curp, _ := strconv.ParseFloat(ss[i][3], 64)
		if curp > maxv {
			maxv = curp
		}
		if curp < minv {
			minv = curp
		}
	}
	fmt.Println(buf.String())
	fmt.Printf("最大：%.2f", maxv)
	fmt.Printf("最小：%.2f", minv)
	fmt.Printf("今日：%.2f", todayp)

}