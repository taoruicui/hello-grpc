package v1

import (
	"context"

	"github.com/pkg/errors"
	"github.com/taoruicui/hello-grpc/pkg/api/v1"
)

var (
	hellos = make(map[int64]*v1.Hello)
)

type HelloServer struct {
}

func NewHelloServer() v1.HelloServiceServer {
	return &HelloServer{}
}

func (hs *HelloServer) Create(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {

	hellos[req.Hello.Id] = req.Hello

	return &v1.HelloResponse{
		Api: "v1",
		Id:  req.Hello.Id,
	}, nil

}

func (hs *HelloServer) Read(ctx context.Context, req *v1.HelloRequest) (*v1.Hello, error) {

	hello, ok := hellos[req.Hello.Id]
	if !ok {
		return nil, errors.Errorf("requested Hello not found given Id: %d", req.Hello.Id)
	}

	return hello, nil
}

func (hs *HelloServer) Update(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {

	hellos[req.Hello.Id] = req.Hello

	return &v1.HelloResponse{
		Api: "v1",
		Id:  req.Hello.Id,
	}, nil
}

func (hs *HelloServer) Delete(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {

	delete(hellos, req.Hello.Id)

	return &v1.HelloResponse{
		Api: "v1",
		Id:  req.Hello.Id,
	}, nil
}
