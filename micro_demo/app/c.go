package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro"
    "micro_demo/pb"
)

func main() {
    // 定义服务，可以传入其他可选参数
    service := micro.NewService(
        micro.Name("greeter.client"),
        )
    service.Init()

    // 创建新的客户端
    greeter := pb.NewGreeterService("greeter", service.Client())

    // 调用greeter
    rsp, err := greeter.Hello(context.TODO(), &pb.HelloRequest{Name: "benben_2015"})
    if err != nil {
        fmt.Println("failed to execute a greeter: ", err)
    }

    // 打印响应请求
    fmt.Println(rsp.Greeting)
}