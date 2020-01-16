package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	http.HandleFunc("/zipdownload.zip", zipHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}

func zipHandler(rw http.ResponseWriter, r *http.Request) {
	zipName := "test.zip"
	//设置rw的header信息中的ctontent-type，对于zip可选以下两种
	// rw.Header().Set("Content-Type", "application/octet-stream")
	rw.Header().Set("Content-Type", "application/zip")
	// 设置rw的header信息中的Content-Disposition为attachment类型
	rw.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", zipName))
	// 向rw中写入zip文件
	f, err := os.Open("./android-studio-ide-173.4907809-linux.zip")
	defer f.Close()
	if err == nil {
		for {
			buf := make([]byte, 1024)
			_, err := f.Read(buf)
			if err == nil {
				rw.Write(buf)
				time.Sleep(10 * time.Microsecond)
			} else {
				return
			}

		}
	}

}
