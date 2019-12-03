package service_client03

import (
	"net/rpc"
	"rpc_demo/service03"
)

type HelloServiceClient struct {
	*rpc.Client
}

var HelloServiceInterface = (*HelloServiceClient)(nil)

func DialHelloService(network, address string) (*HelloServiceClient, error) {
	c, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &HelloServiceClient{Client: c}, nil
}

func (p *HelloServiceClient) Hello(request string, reply *string) error {
	return p.Client.Call(service03.HelloServiceName+".Hello", request, reply)
}