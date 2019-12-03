package main

import (
	"fmt"
	"net"
	"strings"
)

// 读取数据
func handleConnection(conn net.Conn) {

	for {
		buf := make([]byte, 1024)
		if _, err := conn.Read(buf); err == nil {
			result := strings.Replace(string(buf), "\n", "", 1)
			fmt.Println(result)
		} else {
			fmt.Println(err)
		}

	}
}

func main() {

	/*
		Listen: 返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络： "tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket"。
	*/
	listener, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	//todo 限速算法
	fmt.Println("server listen success")
	for {
		//等待客户端接入
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// 使用协程
		go handleConnection(conn)
	}
}
