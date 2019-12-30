package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro"
    "micro_demo/pb"
)

type Greeter struct {
}

func (g *Greeter) Hello(ctx context.Context, req *pb.HelloRequest, rsp *pb.HelloResponse) error {
    rsp.Greeting = "hello" + req.Name
    return nil
}

func main() {
    // 创建新的服务，这里可以传入其他选项
    service := micro.NewService(micro.Name("greeter"))

    // 初始化方法会解析命令行标识
    service.Init()

    // 注册处理器
    err := pb.RegisterGreeterHandler(service.Server(), new(Greeter))
    if err != nil {
        fmt.Println("failed to register a greeter handler: ", err)
    }

    // 运行服务
    if err = service.Run(); err != nil {
        fmt.Println("failed to run a service: ", err)
    }
}