package main

import (
	"fmt"
	"log"
	"rpc_demo/service_client03"
)

func main() {

	client, err := service_client03.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply string
	err = client.Hello("hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
