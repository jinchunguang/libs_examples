package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

//处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
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

var (
	// 图片
	reImg = `"(https?://[^"]+?(\.((jpg)|(jpeg)|(png)|(gif)|(ico))))"`
)

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

//测试是否能拿到数据
func myTest() {
	//1.获取页面内容
	pageStr := GetPageStr("http://www.umei.cc/bizhitupian/meinvbizhi/1.htm")
	fmt.Println(pageStr)
	//2.获取图片链接
	GetImg("http://www.umei.cc/bizhitupian/meinvbizhi/1.htm")
}

//图片下载
func TestDownloadImg() {
	ok := DownloadFile("https://ss0.bdstatic.com/70cFuHSh_Q1YnxGkpoWK1HF6hhy/it/u=1483731740,4186543320&fm=26&gp=0.jpg", "1.jpg")
	if ok {
		fmt.Println("下载成功")
	} else {
		fmt.Println("下载失败")
	}
}

//下载
func DownloadFile(url string, filename string) (ok bool) {
	//发请求
	resp, err := http.Get(url)
	if err != nil {
		HandleError(err, "http.Get")
		return
	}
	//关闭资源
	defer resp.Body.Close()
	//读取响应内容
	fBytes, e := ioutil.ReadAll(resp.Body)
	HandleError(e, "ioutil resp.Body")
	//拼接
	filename = "D:/go_work/img/" + filename
	//写入硬盘
	err = ioutil.WriteFile(filename, fBytes, 644)
	HandleError(err, "http.GetWrite")
	if err != nil {
		return false
	} else {
		return true
	}
}

var (
	//存图片链接的数据通道，string
	chanImageUrls chan string
	//监控通道
	chanTask  chan string
	waitGroup sync.WaitGroup
)

func main() {
	//myTest()
	//TestDownloadImg()

	//1.初始化数据通道
	chanImageUrls = make(chan string, 1000000)
	chanTask = make(chan string, 65)

	//2.爬虫协程
	for i := 1; i < 66; i++ {
		waitGroup.Add(1)
		//获取某个页面所有图片链接
		//strconv.Itoa(i)：将整数转为字符串
		go getImgUrls("http://www.umei.cc/bizhitupian/weimeibizhi/" + strconv.Itoa(i) + ".htm")
	}

	//3.任务统计协程
	waitGroup.Add(1)
	go CheckOk()

	//4.下载协程
	//少开几个下载协程，开5个
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		//下载
		go DownloadImg()
	}
	waitGroup.Wait()
}

//爬当前页所有图片链接，并添加到管道
func getImgUrls(url string) {
	//爬当前页所有图片链接
	urls := getImgs(url)
	//添加到管道
	for _, url := range urls {
		chanImageUrls <- url
	}
	//标志当前协程任务完成
	chanTask <- url
	waitGroup.Done()
}

//拿图片链接
func getImgs(url string) (urls []string) {
	//根据url取内容
	pageStr := GetPageStr(url)
	//获取正则对象
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("找到%d条结果:\n", len(results))
	for _, result := range results {
		//fmt.Println(result)
		//fmt.Println(result)
		url := result[1]
		urls = append(urls, url)
	}
	return
}

//监控65个任务是否完成，完成则关闭通道
func CheckOk() {
	//计数
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成爬取任务\n", url)
		count++
		if count == 65 {
			close(chanImageUrls)
			break
		}
	}
	waitGroup.Done()
}

//下载图片
func DownloadImg() {
	for url := range chanImageUrls {
		//得到全路径
		filename := GetFilenameFromUrl(url, "D:/go_work/src/goapp01/07/img/")
		//保存到硬盘
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
}

//拼接文件名
func GetFilenameFromUrl(url string, dirPath string) (filename string) {
	//strings包的方法，截取最后一个/
	lastIndex := strings.LastIndex(url, "/")
	filename = url[lastIndex+1:]
	//加一个时间戳，防止重名
	timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePrefix + "_" + filename
	filename = dirPath + filename
	return
}
