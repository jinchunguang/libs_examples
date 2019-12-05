package main

import (
    "net/http"
    "time"
)

// 保存 Topic，没有考虑并发问题
var TopicCache = make([]*Topic, 0, 16)

type Topic struct {
    Id        int       `json:"id"`
    Title     string    `json:"title"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
func main() {
    http.HandleFunc("/topic/", handleRequest)
    http.ListenAndServe(":2017", nil)
}
