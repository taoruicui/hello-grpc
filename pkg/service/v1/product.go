package v1

import (
	"context"

	"github.com/pkg/errors"
	"github.com/taoruicui/hello-grpc/pkg/api/v1"
)

var (
	hellos = make(map[int64]*v1.Hello)
)

type ServiceServer struct {
}

func (hs *ServiceServer) Create(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {

	hellos[req.Hello.Id] = req.Hello

	return &v1.HelloResponse{
		Api: "v1",
		Id:  req.Hello.Id,
	}, nil

}

func (hs *ServiceServer) Read(ctx context.Context, req *v1.HelloRequest) (*v1.Hello, error) {
	hello, ok := hellos[req.Hello.Id]
	if !ok {
		return nil, errors.Errorf("requested Hello not found given Id: %d", req.Hello.Id)
	}

	return hello, nil
}

func (hs *ServiceServer) Update(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {

	hellos[req.Hello.Id] = req.Hello

	return &v1.HelloResponse{
		Api: "v1",
		Id:  req.Hello.Id,
	}, nil
}

func (hs *ServiceServer) Delete(ctx context.Context, req *v1.HelloRequest) (*v1.HelloResponse, error) {

	delete(hellos, req.Hello.Id)

	return &v1.HelloResponse{
		Api: "v1",
		Id:  req.Hello.Id,
	}, nil
}

func (hs *ServiceServer) ReadAll(ctx context.Context, req *v1.HelloRequest) (*v1.Hellos, error) {
	res := []*v1.Hello{}
	for _, v := range hellos {
		res = append(res, v)
	}
	return &v1.Hellos{Hellos: res}, nil
}
