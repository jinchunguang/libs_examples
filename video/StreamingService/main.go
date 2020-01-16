package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func videoHandle(w http.ResponseWriter, r *http.Request) {
	file := "oss.mp4"
	video, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer video.Close()
	w.Header().Set("Content-Type", "video/mp4")
	http.ServeContent(w, r, "oss.mp4", time.Now(), video)
}

// 流媒体服务
func main() {
	http.HandleFunc("/video", videoHandle)
	http.ListenAndServe(":10001", nil)
}
