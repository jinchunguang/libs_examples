package main

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	)

func main()  {
	// 读取文件
	cfg, _ := goconfig.LoadConfigFile("conf.ini")
	fmt.Println(cfg)

	// 读取key
	value, _ := cfg.GetValue("database", "host")
	fmt.Println(value)

	// 读取注释
	comment := cfg.GetSectionComments("database")
	fmt.Println(comment)

	// 读取db配置
	sec, _ := cfg.GetSection("database")
	fmt.Println(sec["host"])
	fmt.Println(sec["password"])
	fmt.Println(sec["port"])
	fmt.Println(sec["user"])
}
