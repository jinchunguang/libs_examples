package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc_demo/service02"
)

func main() {

	// 将HelloService类型的对象注册为一个RPC服务
	// rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数
	// 所有注册的方法会放在“HelloService”服务空间之下
	rpc.RegisterName("HelloService", new(service02.HelloService))

	// 建立一个唯一的TCP链接
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// 在该TCP链接上为对方提供RPC服务
		rpc.ServeConn(conn)

		defer conn.Close()
	}
}
