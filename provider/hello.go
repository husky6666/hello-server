package provider

import (
	"context"
	"hello-server/pb_service"
	"log"
)

type HelloProvider struct {
}

func (p *HelloProvider) SayHello(ctx context.Context, request *pb_service.HelloRequest) (*pb_service.HelloReply, error) {
	log.Println("用户名:", request.GetName())

	return &pb_service.HelloReply{
		Message: "你好呀! " + request.GetName(),
	}, nil
}

func NewHelloProvider() *HelloProvider {
	return new(HelloProvider)
}
