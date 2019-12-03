package main

import (
	"fmt"
	"log"

	"io/ioutil"
	"net/http"
	"strings"
)

// 钉钉机器人webhook
var url = "https://oapi.dingtalk.com/robot/send?access_token=daec7a952d07b22295690880e6af06c2412dcc2ca319429bbf93c349c22cb76c"
var textJsonStr, linkJsonStr, markdownJsonStr, actionCardJsonStr, actionCardBtnJsonStr, feedCardJsonStr = "", "", "", "", "", ""

/*
	msgtype 消息类型，此时固定为：text
	content 消息内容
	atMobiles  被@人的手机号(在content里添加@人的手机号)
	isAtAll @所有人时：true，否则为：false
*/
func main() {


	// text类型
	textJsonStr = `{
        "msgtype": "text", 
        "text": {
            "content": "我就是我, 是不一样的烟火@156xxxx8827"
        }, 
        "at": {
            "atMobiles": [
                "156xxxx8827", 
                "189xxxx8325"
            ], 
            "isAtAll": false
        }
    }
    `
	// link类型
	linkJsonStr = `{
        "msgtype": "link", 
        "link": {
            "text": "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。
    而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？", 
            "title": "时代的火车向前开", 
            "picUrl": "", 
            "messageUrl": "https://www.dingtalk.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
        }
    }`

	// markdown类型
	text := "#### 杭州天气 @156xxxx8827\n" +
		"> 9度，西北风1级，空气良89，相对温度73%\n\n" +
		"> ![screenshot](https://gw.alicdn.com/tfs/TB1ut3xxbsrBKNjSZFpXXcXhFXa-846-786.png)\n" +
		"> ###### 10点20分发布 [天气](http://www.thinkpage.cn/) \n"
	markdownJsonStr = `{
        "msgtype": "markdown",
        "markdown": {
            "title":"杭州天气",
            "text": "` + text + `"
        },
       "at": {
           "atMobiles": [
               "156xxxx8827", 
                "189xxxx8325"
           ], 
           "isAtAll": false
       }
    }`

	// 整体跳转ActionCard类型
	actionCardJsonStr = `{
        "actionCard": {
            "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身", 
            "text": "![screenshot](@lADOpwk3K80C0M0FoA) 
    ### 乔布斯 20 年前想打造的苹果咖啡厅 
    Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划", 
            "hideAvatar": "0", 
            "btnOrientation": "0", 
            "singleTitle" : "阅读全文",
            "singleURL" : "https://www.dingtalk.com/"
        }, 
        "msgtype": "actionCard"
    }`

	// 独立跳转ActionCard类型
	actionCardBtnJsonStr = `{
        "actionCard": {
            "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身", 
            "text": "![screenshot](@lADOpwk3K80C0M0FoA) 
    ### 乔布斯 20 年前想打造的苹果咖啡厅 
    Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划", 
            "hideAvatar": "0", 
            "btnOrientation": "0", 
            "btns": [
                {
                    "title": "内容不错", 
                    "actionURL": "https://www.dingtalk.com/"
                }, 
                {
                    "title": "不感兴趣", 
                    "actionURL": "https://www.dingtalk.com/"
                }
            ]
        }, 
        "msgtype": "actionCard"
    }`

	// FeedCard类型
	feedCardJsonStr = `{
			"feedCard": {
				"links": [
						{
							"title": "时代的火车向前开",
							"messageURL": "https://www.dingtalk.com",
							"picURL": "https://mmbiz.qpic.cn/mmbiz_gif/fJjL0RYiaMfeH49RI08jDK3SHswbokXvG6mFQ8CEUousZTFtDJsRT2jWRb5dkJ2TicBywoRqNnLxZJjVelRP4qgg/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1"
						},
						{
							"title": "时代的火车向前开2",
							"messageURL": "https://www.dingtalk.com",
							"picURL": "https://mmbiz.qpic.cn/mmbiz_gif/fJjL0RYiaMfeH49RI08jDK3SHswbokXvG6mFQ8CEUousZTFtDJsRT2jWRb5dkJ2TicBywoRqNnLxZJjVelRP4qgg/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1"
						},
						{
							"title": "时代的火车向前开3",
							"messageURL": "https://www.dingtalk.com",
							"picURL": "https://mmbiz.qpic.cn/mmbiz_gif/fJjL0RYiaMfeH49RI08jDK3SHswbokXvG6mFQ8CEUousZTFtDJsRT2jWRb5dkJ2TicBywoRqNnLxZJjVelRP4qgg/640?wx_fmt=gif&tp=webp&wxfrom=5&wx_lazy=1"
						}
				]
		},
		"msgtype": "feedCard"
		}`

	client := &http.Client{}
	// textJsonStr, linkJsonStr, markdownJsonStr, actionCardJsonStr, actionCardBtnJsonStr, feedCardJsonStr
	req, err := http.NewRequest("POST", url, strings.NewReader(markdownJsonStr))
	if err != nil {
		log.Println(err)
		return
	}
	// header
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))

}
