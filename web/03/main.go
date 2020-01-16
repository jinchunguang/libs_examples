package main

import (
	"log"
	"net/http"
	"web/web"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(wr, r)
	})
}

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

func main() {
	r = webRouter.NewRouter()
	log.Println(r)
	//r.Use(logger)
	// r.Use(timeout)
	// r.Use(ratelimit)
	r.Add("/hello", hello)
	err := http.ListenAndServe(":9007", nil)
	if err != nil {
		log.Fatal(err)
	}
}
