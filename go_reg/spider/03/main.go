package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (

	reImg = `"https?://[^"]+?\.jpg"`
	waitGroup sync.WaitGroup
	downDir = "D:/go_work/img/"
)

func main() {

	//2.爬虫协程
	for i := 1; i < 23000; i++ {
		waitGroup.Add(1)
		// Request the HTML page.
		res, err := http.Get("http://www.netbian.com/desk/23178.htm")
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		// Find the review items
		doc.Find("#main .action").Each(func(i int, s *goquery.Selection) {
			bText := s.Text()
			body, err := GbkToUtf8([]byte(bText))
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(string(body))
			}
			if !strings.Contains(string(body), `壁纸 > 美女壁纸`) {
				return
			}
		})

		imgUrl:=""
		// Find the review items
		doc.Find("#main .endpage .pic p a").Each(func(i int, s *goquery.Selection) {
			bText ,_:= s.Html()
			//获取正则对象
			re := regexp.MustCompile(reImg)
			results := re.FindAllString(bText, -1)
			imgUrl=results[0]
		})

		u, err := url.Parse(imgUrl)
		if err == nil {
			fmt.Println(err.Error())
		}

		fmt.Println(u.Host)
		realImgUrl := strings.Replace(imgUrl, "http://", "//", -1)

		fmt.Println(realImgUrl)
		return
		////发请求
		r, err := http.Get(realImgUrl)
		if err!=nil {
			fmt.Printf(err.Error())
		}

		fmt.Println(r.Status)
		return
		////关闭资源
		//defer respImg.Body.Close()
		////读取响应内容
		//fBytes, _ := ioutil.ReadAll(respImg.Body)
		//
		////拼接
		////filename = "D:/go_work/img/" + strconv.Itoa(i)
		//lastIndex := strings.LastIndex(imgUrl, "/")
		//filename := imgUrl[lastIndex+1:]
		////加一个时间戳，防止重名
		//timePrefix := strconv.Itoa(int(time.Now().UnixNano()))
		//filename = timePrefix + "_" + filename
		//filename = "D:/go_work/img/" + filename
		//
		////写入硬盘
		//err = ioutil.WriteFile(filename, fBytes, 644)
		//if err!=nil {
		//	fmt.Printf("%s 下载失败\n", filename)
		//} else {
		//
		//	fmt.Printf("%s 下载成功\n", filename)
		//}
		return

		os.Exit(0)
		time.Sleep(3)
		////得到全路径
		//filename := GetFilenameFromUrl("", "D:/go_work/src/goapp01/07/img/")
		////保存到硬盘
		//ok := DownloadFile(url, filename)
		//if ok {
		//	fmt.Printf("%s 下载成功\n", filename)
		//} else {
		//	fmt.Printf("%s 下载失败\n", filename)
		//}

		waitGroup.Done()
	}

	waitGroup.Wait()
	//6.爬图片链接
	//GetImg("http://www.netbian.com/desk/1.htm")
}

//处理异常
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
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
	filename = downDir + filename
	//写入硬盘
	err = ioutil.WriteFile(filename, fBytes, 644)
	HandleError(err, "http.GetWrite")
	if err != nil {
		return false
	} else {
		return true
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

//func GetLink(url string)  {
//	pageStr := GetPageStr(url)
//	fmt.Println(pageStr)
//	re := regexp.MustCompile(reLink)
//	results := re.FindAllStringSubmatch(pageStr, -1)
//	fmt.Printf("找到%d条结果:\n",len(results))
//	for _, result := range results {
//		//fmt.Println(result)
//		fmt.Println(result[1])
//	}
//}
//func GetImg(url string)  {
//	pageStr := GetPageStr(url)
//	//fmt.Println(pageStr)
//	re := regexp.MustCompile(reImg)
//	results := re.FindAllStringSubmatch(pageStr, -1)
//	fmt.Printf("找到%d条结果:\n",len(results))
//	for _, result := range results {
//		//fmt.Println(result)
//		fmt.Println(result[0])
//	}
//}
