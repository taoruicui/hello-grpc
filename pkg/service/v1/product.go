package v1

import (
	"context"

	"time"

	"github.com/pkg/errors"
	"github.com/taoruicui/hello-grpc/pkg/api/v1"
	"github.com/taoruicui/hello-grpc/pkg/metrics"
)

var (
	products = make(map[int64]*v1.Product)
)

type ServiceServer struct {
}

// Wrapper code for plugin metrics
func (hs *ServiceServer) CreateProduct(ctx context.Context, req *v1.ProductRequest) (res *v1.ProductResponse, err error) {
	metrics.PromMetrics.APIHits.WithLabelValues("CreateProduct").Inc()
	start := time.Now()
	defer func() {
		metrics.PromMetrics.APILatency.WithLabelValues("CreateProduct").Observe(float64(time.Now().Sub(start)))
		if err != nil {
			metrics.PromMetrics.APIErrors.WithLabelValues("CreateProduct").Inc()
		}
	}()
	return createProduct(ctx, req)
}

func (hs *ServiceServer) ReadProduct(ctx context.Context, req *v1.ProductRequest) (res *v1.Product, err error) {
	metrics.PromMetrics.APIHits.WithLabelValues("ReadProduct").Inc()
	start := time.Now()
	defer func() {
		metrics.PromMetrics.APILatency.WithLabelValues("ReadProduct").Observe(float64(time.Now().Sub(start)))
		if err != nil {
			metrics.PromMetrics.APIErrors.WithLabelValues("ReadProduct").Inc()
		}
	}()
	return readProduct(ctx, req)
}

func (hs *ServiceServer) UpdateProduct(ctx context.Context, req *v1.ProductRequest) (res *v1.ProductResponse, err error) {
	metrics.PromMetrics.APIHits.WithLabelValues("UpdateProduct").Inc()
	start := time.Now()
	defer func() {
		metrics.PromMetrics.APILatency.WithLabelValues("UpdateProduct").Observe(float64(time.Now().Sub(start)))
		if err != nil {
			metrics.PromMetrics.APIErrors.WithLabelValues("UpdateProduct").Inc()
		}
	}()
	return updateProduct(ctx, req)
}

func (hs *ServiceServer) DeleteProduct(ctx context.Context, req *v1.ProductRequest) (res *v1.ProductResponse, err error) {
	metrics.PromMetrics.APIHits.WithLabelValues("DeleteProduct").Inc()
	start := time.Now()
	defer func() {
		metrics.PromMetrics.APILatency.WithLabelValues("DeleteProduct").Observe(float64(time.Now().Sub(start)))
		if err != nil {
			metrics.PromMetrics.APIErrors.WithLabelValues("DeleteProduct").Inc()
		}
	}()
	return deleteProduct(ctx, req)
}

func (hs *ServiceServer) ReadAllProducts(ctx context.Context, req *v1.ProductRequest) (res *v1.Products, err error) {
	metrics.PromMetrics.APIHits.WithLabelValues("ReadAllProducts").Inc()
	start := time.Now()
	defer func() {
		metrics.PromMetrics.APILatency.WithLabelValues("ReadAllProducts").Observe(float64(time.Now().Sub(start)))
		if err != nil {
			metrics.PromMetrics.APIErrors.WithLabelValues("ReadAllProducts").Inc()
		}
	}()
	return readAllProducts(ctx, req)
}

func createProduct(ctx context.Context, req *v1.ProductRequest) (*v1.ProductResponse, error) {
	products[req.Product.Id] = req.Product

	return &v1.ProductResponse{
		Api: "v1",
		Id:  req.Product.Id,
	}, nil
}

func readProduct(ctx context.Context, req *v1.ProductRequest) (*v1.Product, error) {
	product, ok := products[req.Product.Id]
	if !ok {
		return nil, errors.Errorf("requested Hello not found given Id: %d", req.Product.Id)
	}

	return product, nil
}

func updateProduct(ctx context.Context, req *v1.ProductRequest) (*v1.ProductResponse, error) {
	products[req.Product.Id] = req.Product

	return &v1.ProductResponse{
		Api: "v1",
		Id:  req.Product.Id,
	}, nil
}

func deleteProduct(ctx context.Context, req *v1.ProductRequest) (*v1.ProductResponse, error) {

	delete(products, req.Product.Id)

	return &v1.ProductResponse{
		Api: "v1",
		Id:  req.Product.Id,
	}, nil
}

func readAllProducts(ctx context.Context, req *v1.ProductRequest) (*v1.Products, error) {
	var res []*v1.Product
	for _, v := range products {
		res = append(res, v)
	}
	return &v1.Products{Products: res}, nil
}
