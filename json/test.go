package main

import (
	"encoding/json"
	"fmt"
)

func sayHi(v ...interface{}) {
	ret, _ := json.Marshal(v)
	fmt.Println(string(ret))
}

func main() {
	sayHi(100)
	sayHi(1000.00)
	sayHi("hello world")
	sayHi(map[string]string{"name": "张三", "age": "unknown"})
	sayHi([5]float32{1000.0, 2.0, 3.4, 7.0, 50.0})
}
