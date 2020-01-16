package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	udpAddr, _ := net.ResolveUDPAddr("udp4", "127.0.0.1:9998")

	//连接udpAddr，返回 udpConn
	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println("udp dial ok ")

	// 发送数据
	len, err := udpConn.Write([]byte("上报日志文件:xxxxxx\r\n"))
	if err != nil {
		return
	}
	fmt.Println("client write len:", len)

	//读取数据
	buf := make([]byte, 1024)
	len, _ = udpConn.Read(buf)
	fmt.Println("client read len:", len)
	fmt.Println("client read data:", string(buf))

}
