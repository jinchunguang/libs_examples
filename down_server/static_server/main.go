package main

import (
	"net/http"
)

/**
静态文件服务
 */
func main() {
	fs := http.FileServer(http.Dir("./tmp"))
	http.Handle("/download/", http.StripPrefix("/download/", fs))
	http.ListenAndServe(":8080", nil)
}
