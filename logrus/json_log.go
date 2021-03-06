package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

// logrus提供了New()函数来创建一个logrus的实例。
// 项目中，可以创建任意数量的logrus实例。

func main() {
	var log = logrus.New()
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	log.Out = os.Stdout

	// 为当前logrus实例设置消息输出格式为json格式。
	// 同样地，也可以单独为某个logrus实例设置日志级别和hook，这里不详细叙述。
	log.Formatter = &logrus.JSONFormatter{}

	log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.Debug("调试信息")
	logrus.Info("自定义的信息")
	logrus.Warn("警告信息")
	logrus.Error("一般错误信息")
	logrus.Fatal("严重错误信息") //log之后会调用os.Exit(1)
	logrus.Panic("系统级别错误") //log之后会panic()
}
