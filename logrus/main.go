package main

import (
	logrus "github.com/sirupsen/logrus"
	"os"
)
func main()  {

	// 设置日志格式为json格式　自带的只有两种样式logrus.JSONFormatter{}和logrus.TextFormatter{}
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 设置将日志输出到标准输出（默认的输出为stderr，标准错误）
	// 日志消息输出可以是任意的io.writer类型
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	//logrus.SetLevel(logrus.WarnLevel)


	entry := logrus.WithFields(logrus.Fields{"request_id": 1001, "user_ip": "255.255.276.90"})
	entry.Info("something happened on that request")
	entry.Warn("something not great happened")

	//logrus.Debug("调试信息")
	//logrus.Info("自定义的信息")
	//logrus.Warn("警告信息")
	//logrus.Error("一般错误信息")
	//logrus.Fatal("严重错误信息")   //log之后会调用os.Exit(1)
	//logrus.Panic("系统级别错误")   //log之后会panic()




}