/**
 * @Author : jinchunguang
 * @Date : 19-10-31 下午12:35
 * @Project : sty
 */
package main

import (
    "expvar"
    "fmt"
    "net/http"
)

var visits = expvar.NewInt("visits")

func handler(w http.ResponseWriter, r *http.Request) {
    visits.Add(1)
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8999", nil)
}

