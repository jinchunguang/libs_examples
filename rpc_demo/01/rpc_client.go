package _1

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {


	// 拨号RPC服务
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	// 调用client.Call时，第一个参数是用点号链接的RPC服务名字和方法名字，第二和第三个参数分别我们定义RPC方法的两个参数。
	err = client.Call("HelloService.Hello", "world", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}