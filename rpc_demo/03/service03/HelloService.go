package service03

import "net/rpc"

type HelloService struct {}

// 1 服务的名字
const HelloServiceName = "HelloService"

// 2 服务要实现的详细方法列表
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello " + request
	return nil
}

type HelloServiceInterface = interface {
	Hello(request string, reply *string) error
}

// 3 注册该类型服务的函数
func RegisterHelloService(svc HelloServiceInterface) error {
	return rpc.RegisterName(HelloServiceName, svc)
}