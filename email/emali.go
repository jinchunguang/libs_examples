/**
 * @Author : jinchunguang
 * @Date : 19-10-18 下午1:35
 * @Project : learning
 */
package main

import (
    "github.com/go-gomail/gomail"
    "log"
    "strings"
)

type EmailParam struct {
    // 邮箱服务器地址
    ServerHost string
    // 邮箱服务器端口
    ServerPort int
    // 发件人邮箱地址
    FromEmail string
    // 发件人邮箱密码
    FromPasswd string
    // 接收者邮件，多个则以英文逗号(“,”)隔开，不能为空
    Toers string
    // 抄送者邮件，多个则以英文逗号(“,”)隔开，可以为空
    CCers string
}

// 邮件服务器、发件人账号、密码(可以从配置文件读取)
var serverHost, fromEmail, fromPasswd string
var serverPort int
var m *gomail.Message

func InitEmail(ep *EmailParam) {
    toers := []string{}

    // smtp.exmail.qq.com(使用SSL，端口号465)
    serverHost = "smtp.exmail.qq.com"
    serverPort = 465
    fromEmail = ""
    fromPasswd = ""

    m = gomail.NewMessage()

    if len(ep.Toers) == 0 {
        return
    }

    for _, tmp := range strings.Split(ep.Toers, ",") {
        toers = append(toers, strings.TrimSpace(tmp))
    }

    // 多个收件人，分割
    m.SetHeader("To", toers...)

    // 抄送列表
    if len(ep.CCers) != 0 {
        for _, tmp := range strings.Split(ep.CCers, ",") {
            toers = append(toers, strings.TrimSpace(tmp))
        }
        m.SetHeader("Cc", toers...)
    }

    // 发件人
    // 第三个参数为发件人别名，如"李大锤"，可以为空
    m.SetAddressHeader("From", fromEmail, "")
}

// body支持html格式字符串
func SendEmail(subject, body string) {
    // 主题
    m.SetHeader("Subject", subject)

    // 正文
    m.SetBody("text/html", body)

    d := gomail.NewDialer(serverHost, serverPort, fromEmail, fromPasswd)
    // 发送
    err := d.DialAndSend(m)
    if err != nil {
        panic(err)
    }
}

func main() {

    myToers := "" // 逗号隔开
    myCCers := ""                           // "readchy@163.com"

    subject := "这是主题"
    body := `这是正文<br>
            <h3>这是标题</h3>
             Hello <a href = "http://www.latelee.org">主页</a><br>`
    // 结构体赋值
    myEmail := &EmailParam{
        ServerHost: serverHost,
        ServerPort: serverPort,
        FromEmail:  fromEmail,
        FromPasswd: fromPasswd,
        Toers:      myToers,
        CCers:      myCCers,
    }
    log.Println("init email.\n")
    InitEmail(myEmail)
    SendEmail(subject, body)
}
