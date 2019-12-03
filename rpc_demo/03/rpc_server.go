package main

import (
	"log"
	"net"
	"net/rpc"
	"rpc_demo/service03"
)

func main() {
	service03.RegisterHelloService(new(service03.HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}
