package v1

import (
	"context"

	"github.com/pkg/errors"
	"github.com/taoruicui/hello-grpc/pkg/api/v1"
)

var (
	products = make(map[int64]*v1.Product)
)

type ServiceServer struct {
}

func (hs *ServiceServer) Create(ctx context.Context, req *v1.ProductRequest) (*v1.ProductResponse, error) {

	products[req.Product.Id] = req.Product

	return &v1.ProductResponse{
		Api: "v1",
		Id:  req.Product.Id,
	}, nil

}

func (hs *ServiceServer) Read(ctx context.Context, req *v1.ProductRequest) (*v1.Product, error) {
	product, ok := products[req.Product.Id]
	if !ok {
		return nil, errors.Errorf("requested Hello not found given Id: %d", req.Product.Id)
	}

	return product, nil
}

func (hs *ServiceServer) Update(ctx context.Context, req *v1.ProductRequest) (*v1.ProductResponse, error) {

	products[req.Product.Id] = req.Product

	return &v1.ProductResponse{
		Api: "v1",
		Id:  req.Product.Id,
	}, nil
}

func (hs *ServiceServer) Delete(ctx context.Context, req *v1.ProductRequest) (*v1.ProductResponse, error) {

	delete(products, req.Product.Id)

	return &v1.ProductResponse{
		Api: "v1",
		Id:  req.Product.Id,
	}, nil
}

func (hs *ServiceServer) ReadAll(ctx context.Context, req *v1.ProductRequest) (*v1.Products, error) {
	var res []*v1.Product
	for _, v := range products {
		res = append(res, v)
	}
	return &v1.Products{Products: res}, nil
}
