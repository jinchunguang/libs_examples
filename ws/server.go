package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	// 设置缓冲
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 协商压缩
	//EnableCompression:true,
	// 跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// 调用连接的WriteMessage和ReadMessage方法以一片字节发送和接收消息。
	// p是一个[]字节，messageType是一个值为websocket.BinaryMessage或websocket.TextMessage的int
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}

		for {
			conn.WriteMessage(messageType, []byte("[Server]当前时间:"+time.Now().Format("2006-01-02 15:04:05")))
			time.Sleep(5 * time.Second)
		}
	}

	// 使用io.WriteCloser和io.Reader接口发送和接收消息。
	// 发送消息，请调用连接NextWriter方法以获取io.WriteCloser，将消息写入writer并在完成后关闭writer。
	// 接收消息，请调用连接NextReader方法以获取io.Reader并读取，直到返回io.EOF。
	//for {
	//
	//	messageType, r, err := conn.NextReader()
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//
	//	log.Println(&conn)
	//	w, err := conn.NextWriter(messageType)
	//	if err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	if _, err := io.Copy(w, r); err != nil {
	//		log.Println(err)
	//		return
	//	}
	//	if err := w.Close(); err != nil {
	//		log.Println(err)
	//		return
	//	}
	//}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":2000", nil)
}
