package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		// next handler
		next.ServeHTTP(wr, r)
		timeElapsed := time.Since(timeStart)
		logger.Println(timeElapsed)
	})
}

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func main() {
	http.Handle("/hello", timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":9007", nil)
	if err != nil {
		log.Fatal(err)
	}
}
