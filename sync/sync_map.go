package main

import (
	"fmt"
	"sync"
)

// 线程不安全
func main() {

	var scene sync.Map

	// 将键值对保存到sync.Map
	scene.Store("zhansan", 97)
	scene.Store("lisi", 100)
	scene.Store("wangwu", 200)

	// 从sync.Map中根据键取值
	fmt.Println(scene.Load("london"))
	value, _ := scene.Load("zhansan")
	fmt.Println(value.(int))

	// 根据键删除对应的键值对
	scene.Delete("london")

	// 遍历所有sync.Map中的键值对
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
