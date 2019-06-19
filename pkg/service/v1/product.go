package v1

import (
	"github.com/taoruicui/hello-grpc/pkg/api/v1"
	"context"
	"fmt"
)

type HelloServer struct {

}

func NewHelloServer() v1.HelloServiceServer {
	return &HelloServer{}
}

func (hs *HelloServer) Create(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {
	fmt.Println(req.GetHello())
	return nil, nil
}
