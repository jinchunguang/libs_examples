package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {

	//基本示例
	logrus.WithFields(logrus.Fields{
		"animal": "dog",
	}).Info("一条舔狗出现了。")

	// 进阶示例
	// 创建一个新的logger实例。可以创建任意多个。
	var log = logrus.New()
	// 会记录info及以上级别 (warn, error, fatal, panic)
	// log.SetLevel(logrus.WarnLevel)
	// 设置日志输出为os.Stdout
	log.Out = os.Stdout
	// 可以设置像文件等任意`io.Writer`类型作为日志输出
	// file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	// if err == nil {
	//  log.Out = file
	// } else {
	//  log.Info("Failed to log to file, using default stderr")
	// }

	log.WithFields(logrus.Fields{
		"animal": "dog",
		"size":   10,
	}).Info("一群舔狗出现了。")

	// Logrus有七个日志级别：Trace, Debug, Info, Warning, Error, Fataland Panic。
	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	// log.Error("Something failed but I'm not quitting.")
	// // 记完日志后会调用os.Exit(1)
	// log.Fatal("Bye.")
	// // 记完日志后会调用 panic()
	// log.Panic("I'm bailing.")

	// 字段
	log.WithFields(logrus.Fields{
		"event": "event",
		"topic": "topic",
		"key":   "key",
	}).Info("Failed to send event")

	// 默认字段
	/*
	   通常，将一些字段始终附加到应用程序的全部或部分的日志语句中会很有帮助。例如，你可能希望始终在请求的上下文中记录request_id和user_ip。
	   区别于在每一行日志中写上log.WithFields(log.Fields{"request_id": request_id, "user_ip": user_ip})
	*/
	requestLogger := logrus.WithFields(logrus.Fields{"request_id": time.Now().UnixNano(), "user_ip": "127.0.0.1"})
	requestLogger.Info("something happened on that request")
	requestLogger.Warn("something not great happened")

	/*
	   日志条目
	   除了使用WithField或WithFields添加的字段外，一些字段会自动添加到所有日志记录事中:

	   time：记录日志时的时间戳
	   msg：记录的日志信息
	   level：记录的日志级别
	*/

	/*
	   Hooks
	   你可以添加日志级别的钩子（Hook）。例如，向异常跟踪服务发送Error、Fatal和Panic、信息到StatsD或同时将日志发送到多个位置，例如syslog。
	*/

}
