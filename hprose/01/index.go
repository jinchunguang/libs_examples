package main

import (
	"github.com/hprose/hprose-golang/rpc"
	"net/http"
)

func hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	service := rpc.NewHTTPService()
	service.AddFunction("hello", hello)
	http.ListenAndServe(":7070", service)
}
