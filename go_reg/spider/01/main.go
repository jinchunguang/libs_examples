package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

var (
	// qq邮箱
	reQQEmail = `(\d+)@qq.com`

	// 匹配邮箱
	reEmail = `\w+@\w+\.\w+(\.\w+)?`

	// 链接
	reLink = `href="(https?://[\s\S]+?)"`

	// 手机号
	rePhone=`1[3456789]\d\s?\d{4}\s?\d{4}`

	// 身份证
	reIdCard=`[12345678]\d{5}((19\d{2})|(20[01]))((0[1-9]|[1[012]]))((0[1-9])|[12]\d|[3[01]])\d{3}[\dXx]`

	// 图片
	reImg=`"(https?://[^"]+?(\.((jpg)|(jpeg)|(png)|(gif)|(ico))))"`
)

func main() {
	//1.爬邮箱
	//GetEmail()
	//2.抽取爬邮箱的方法
	//GetEmail2("http://tieba.baidu.com/p/2544042204")
	//3.爬超链接
	//GetLink("http://www.baidu.com/s?wd=岛国%20留下邮箱")
	//4.爬手机号
	//GetPhone("http://www.zhaohaowang.com/")
	//5.爬身份证
	//GetIdcard("http://henan.qq.com/a/20171107/069413.htm")
	//6.爬图片链接
	GetImg("http://image.baidu.com/search/index?tn=baiduimage&ps=1&ct=201326592&lm=-1&cl=2&nc=1&ie=utf-8&word=%E7%BE%8E%E5%A5%B3")
}

//爬邮箱
func GetEmail() {
	//1.发送http请求，获取页面内容
	resp, err := http.Get("http://tieba.baidu.com/p/2544042204")
	//处理异常
	HandleError(err, "http.Get url")
	//关闭资源
	defer resp.Body.Close()
	//接收页面
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	//打印页面内容
	pageStr := string(pageBytes)
	fmt.Println(pageStr)

	//2.捕获邮箱，先搞定qq邮箱
	//传入正则
	re := regexp.MustCompile(reQQEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		//fmt.Println(result)
		fmt.Printf("email=%s qq=%s\n", result[0], result[1])
	}
}

//处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

//抽取的爬邮箱的方法
func GetEmail2(url string) {
	//爬页面所有数据
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reEmail)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		fmt.Println(result)
	}
}

//根据url获取页面内容
func GetPageStr(url string) (pageStr string) {
	//1.发送http请求，获取页面内容
	resp, err := http.Get(url)
	//处理异常
	HandleError(err, "http.Get url")
	//关闭资源
	defer resp.Body.Close()
	//接收页面
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	//打印页面内容
	pageStr = string(pageBytes)
	return pageStr
}

func GetLink(url string)  {
	pageStr := GetPageStr(url)
	fmt.Println(pageStr)
	re := regexp.MustCompile(reLink)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("找到%d条结果:\n",len(results))
	for _, result := range results {
		//fmt.Println(result)
		fmt.Println(result[1])
	}
}

func GetPhone(url string)  {
	pageStr := GetPageStr(url)
	fmt.Println(url)
	re := regexp.MustCompile(rePhone)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("找到%d条结果:\n",len(results))
	for _, result := range results {
		//fmt.Println(result)
		fmt.Println(result)
	}
}

func GetIdcard(url string)  {
	pageStr := GetPageStr(url)
	// fmt.Println(pageStr)
	re := regexp.MustCompile(reIdCard)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("找到%d条结果:\n",len(results))
	for _, result := range results {
		//fmt.Println(result)
		fmt.Println(result)
	}
}

func GetImg(url string)  {
	pageStr := GetPageStr(url)
	//fmt.Println(pageStr)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("找到%d条结果:\n",len(results))
	for _, result := range results {
		//fmt.Println(result)
		fmt.Println(result[0])
	}
}