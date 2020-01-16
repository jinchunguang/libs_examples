/**
 * @Author : jinchunguang
 * @Date : 19-10-13 下午10:49
 * @Project : regexp
 */
package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// 获取手机号
func spiderPhone() {

	// 新增了166、198、199号段的手机号
	phoneRegexp := `0?(13|14|15|17|18)[0-9]{9}`
	url := "https://www.haomagujia.com/"

	// 获取网页内容
	resp, err := http.Get(url)
	handleErr(err)
	html, err := ioutil.ReadAll(resp.Body)

	// 正则匹配
	reg := regexp.MustCompile(phoneRegexp)
	matchReg := reg.FindAllString(string(html), -1)

	// 打印结果
	for _, v := range matchReg {
		fmt.Println(v)
	}
}

// 获取手机号分组
func spiderPhoneGroup() {
	// 新增了166、198、199号段的手机号
	phoneRegexp := `0?(13|14|15|17|18)[0-9]{9}`
	url := "https://www.haomagujia.com/"

	// 获取网页内容
	resp, err := http.Get(url)
	handleErr(err)
	html, err := ioutil.ReadAll(resp.Body)

	// 正则匹配
	reg := regexp.MustCompile(phoneRegexp)
	// 打印10个匹配,-1 获取所有
	matchReg := reg.FindAllStringSubmatch(string(html), 15)

	// 打印结果
	for _, v := range matchReg {
		fmt.Println(v)
	}
}

// 获取邮箱
func spiderEmail() {

	regexpEmail := `\w[-\w.+]*@([A-Za-z0-9][-A-Za-z0-9]+\.)+[A-Za-z]{2,14}`
	url := "http://szb.wsnews.com.cn/pad/qdwk/html/2018-11/29/content_228697.htm"

	// 由于访问了https
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	// 获取网页内容
	resp, err := client.Get(url)
	handleErr(err)
	html, err := ioutil.ReadAll(resp.Body)

	// 正则匹配
	reg := regexp.MustCompile(regexpEmail)
	// -1 获取所有
	matchReg := reg.FindAllString(string(html), -1)

	// 打印结果
	for _, v := range matchReg {
		fmt.Println(v)
	}
}

// 获取超链接
func spiderUrl() {
	// regexpUrl := `<a .*?href="(.*?)".*?>`
	// 其结果可能存在锚点，全路径，相对路径,这里只获取http开头的
	regexpUrl := `<a .*?href="(http.*?)".*?>`
	url := "https://www.hao123.com"

	// 获取网页内容
	resp, err := http.Get(url)
	handleErr(err)
	html, err := ioutil.ReadAll(resp.Body)

	// 正则匹配
	reg := regexp.MustCompile(regexpUrl)
	// -1 获取所有
	matchReg := reg.FindAllStringSubmatch(string(html), -1)

	// 打印结果
	for _, v := range matchReg {
		fmt.Println(v[1])
	}
}

// 获取身份证
func spiderIdentityCard() {

	regexpUrl := `\d{17}[\d|x]|\d{15}`
	// regexpUrl := `\d{6}(18|19|20)?\d{2}(0[1-9]|1[12])(0[1-9]|[12]\d|3[01])\d{3}(\d|X)`
	url := "https://baijiahao.baidu.com/s?id=1594520665945691473&wfr=spider&for=pc"
	// 由于访问了https
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	// 获取网页内容
	resp, err := client.Get(url)
	handleErr(err)
	html, err := ioutil.ReadAll(resp.Body)

	// 正则匹配
	reg := regexp.MustCompile(regexpUrl)
	// -1 获取所有
	matchReg := reg.FindAllStringSubmatch(string(html), -1)

	// 打印结果
	for _, v := range matchReg {
		fmt.Println(v)
	}
}

// 获取图片
func spiderImgUrl() [][]string {

	regexpUrl := `<img[^>]*src[=\'\"\s]+([^\"\']*)[\"\']?[^>]*>`
	url := "https://www.zswxy.cn/articles/16972.html"

	// 获取网页内容
	resp, err := http.Get(url)
	handleErr(err)
	html, err := ioutil.ReadAll(resp.Body)

	// 正则匹配
	reg := regexp.MustCompile(regexpUrl)
	// -1 获取所有
	matchReg := reg.FindAllStringSubmatch(string(html), -1)

	// 打印结果
	for _, v := range matchReg {
		fmt.Println(v)
	}
	return matchReg

}

// 下载图片
func downloadImg(imagesUrls [][]string) {

	var dir string
	var imgName string
	for _, v := range imagesUrls {

		// 获取alt属性
		imgHtml := v[0]
		reg := regexp.MustCompile(`<img[^>]*alt[=\'\"\s]+([^\"\']*)[\"\']?[^>]*>`)
		alts := reg.FindAllStringSubmatch(string(imgHtml), -1)
		for _, alt := range alts {
			if len(alt) > 1 {
				// 使用alt做图片名，并且将空格替换为下划线
				imgName = strings.Replace(alt[1], " ", "_", -1)
			}
		}

		// 下载图片
		resp, err := http.Get(v[1])
		handleErr(err)
		imgBytes, err := ioutil.ReadAll(resp.Body)
		dir = "./images/" + imgName + "_" + strconv.Itoa(int(time.Now().UnixNano())) + ".jpg"
		err = ioutil.WriteFile(dir, imgBytes, 0666)
		if err == nil {
			log.Println("下载成功", dir, v[1])
		} else {
			handleErr(err)
		}

	}
	log.Println("图片总数", len(imagesUrls))
}
func main() {
	// spiderPhone()
	// spiderPhoneGroup()
	// spiderEmail()
	// spiderUrl()
	spiderIdentityCard()
	// spiderImgUrl()
	// downloadImg(spiderImgUrl())
}
