# Go Micro


Go Micro是一个微服务开发框架。Go Micro提供了分布式系统开发的核心要求，包括RPC和事件驱动的通信。

## 特点

自动服务注册与名称解析：服务发现是微服务开发中的核心，用于解析服务名与地址。consul是Go Micro默认的服务发现注册中心。发现系统可插拔，其他插件像etcd，kubernetes，zookeeper。
负载均衡：在服务发现之上构建了负载均衡机制。使用随机处理过的哈希负载均衡机制来保证对服务请求颁布的均匀分布。
消息编码：支持基于内容类型（content-type）动态编码消息。客户端和服务端会一起使用content-type格式来对Go进行无缝编/解码。content-type默认包含proto-rpc和json-rpc。
Request/Response：RPC通信基于支持双向流的请求/响应方式，提供同步通信机制，请求发送到服务时，会自动解析，负载均衡，拨号，转成字节流。
异步消息：发布订阅等功能内置在异步通信与事件驱动架构中，事件通知在微服务开发中处于核心位置。默认的消息传递使用点到点http/1.1，激活tls时则使用http2。
可插拔接口：Go Micro为每个分布式系统抽象出接口，因此，Go Micro的接口都是可插拔的，允许其在运行时不可知的情况下仍可支持。


## 编写服务

### 服务原型

微服务中有个关键需求点，就是接口的强定义。Micro使用protobuf来完成这个需求。以gRPC使用简介文章中的服务为例，定义Greeter处理器，它有一个Hello方法。它有HelloRequest入参对象以及HelloResponse出参对象，两个对象都有一个字符串类型的参数。
helloworld.proto文件：


### proto

```
syntax = "proto3";

package pb;

service Greeter {
    rpc Hello (HelloRequest) returns (HelloResponse) {
    }
}

message HelloRequest {
    string name = 1;
}

message HelloResponse {
    string greeting = 2;
}
```

生成协议

`protoc --proto_path=. --micro_out=. --go_out=. ./helloworld.proto`



执行完命令后，将会在当前目录下生成`helloworld.pb.go`文件和`helloworld.micro.go`文件。
其中`helloworld.micro.go`文件中分别定义了`Greeter`服务的客户端API和服务端API，如下：

### server

```
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
```

### client

```
package main

import (
    "context"
    "fmt"
    "github.com/micro/go-micro"
    "micro_demo/pb"
)

func main() {
    // 定义服务，可以传入其他可选参数
    service := micro.NewService(micro.Name("greeter.client"))
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
```

### 测试

```
✗ go run s.go 
2019-12-30 17:30:25.871326 I | Transport [http] Listening on [::]:39939
2019-12-30 17:30:25.871384 I | Broker [http] Connected to [::]:45933
2019-12-30 17:30:25.872476 I | Registry [mdns] Registering node: greeter-1aff7132-fb94-45ac-86c1-c8c7134a29c1

```

访问:

```
 ✗ go run c.go 
hellobenben_2015
```

---

## 使用consul

在启动go-micro服务时，可以看到使用的服务注册是mdns。这节来了解下如何使用consul。默认的服务发现是在同一台机器上，在生产环境下，服务与服务发现部署到同一机器，这不是个很好的实践。所以把consul机器独立出来就很有必要。


当我们服务越来越多，如果服务配置了弹性伸缩，或者当服务不可用时，我们需要随时动态掌握可以使用的服务数量，并向可提供响应的服务发送请求。这时我们需要服务发现功能，当新增服务时，服务可以自动向consul注册，客户端直接向consul发送请求，获取可用服务的地址和端口；当服务不可用时，动态的更新consul，删除该服务在consul中的列表

 

docker安装consul

docker run --name consul1 -d -p 8500:8500 -p 8300:8300 -p 8301:8301 -p 8302:8302 -p 8600:8600 consul:latest agent -server -bootstrap-expect 2 -ui -bind=0.0.0.0 -client=0.0.0.0

- 8500 http 端口，用于 http 接口和 web ui
- 8300 server rpc 端口，同一数据中心 consul server 之间通过该端口通信
- 8301 serf lan 端口，同一数据中心 consul client 通过该端口通信
- 8302 serf wan 端口，不同数据中心 consul server 通过该端口通信
- 8600 dns 端口，用于服务发现
- -bbostrap-expect 2: 集群至少两台服务器，才能选举集群leader
- -ui：运行 web 控制台
- -bind： 监听网口，0.0.0.0 表示所有网口，如果不指定默认未127.0.0.1，则无法和容器通信
- -client ： 限制某些网口可以访问



✗ docker inspect --format '{{ .NetworkSettings.IPAddress }}' consul1
172.17.0.2



docker run --name consul2 -d -p 8501:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join 172.17.0.2
docker run --name consul3 -d -p 8502:8500 consul agent -server -ui -bind=0.0.0.0 -client=0.0.0.0 -join 172.17.0.2


