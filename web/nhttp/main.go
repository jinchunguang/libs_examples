package main

import (
    "github.com/jinchunguang/demo/web/nhttp/router"
    "log"
    "net/http"
)

// var logger = log.New(os.Stdout, "", 0)

func logger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
        next.ServeHTTP(wr, r)
        // logger.Println("timeMiddleware")
    })
}

func helloHandler(wr http.ResponseWriter, r *http.Request) {
    wr.Write([]byte("hello"))
}

func main() {
    r := router.NewRouter()
    r.Use(logger)
    r.Add("/hello", http.HandlerFunc(helloHandler))
    // http.Handle("/hello", timeMiddleware(http.HandlerFunc(hello)))
    err := http.ListenAndServe(":9007", r)
    if err != nil {
        log.Fatal(err)
    }
}
