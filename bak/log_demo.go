package bak

import (
	log "github.com/sirupsen/logrus"
)


func initLog() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	initLog()
	log.WithFields(log.Fields{
		"age": 12,
		"name":   "xiaoming",
		"sex": 1,
	}).Info("小明来了")

	log.WithFields(log.Fields{
		"age": 13,
		"name":   "xiaohong",
		"sex": 0,
	}).Error("小红来了")

	log.WithFields(log.Fields{
		"age": 14,
		"name":   "xiaofang",
		"sex": 1,
	}).Fatal("小芳来了")
}
