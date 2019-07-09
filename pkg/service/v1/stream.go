package v1

import (
	"io"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/taoruicui/hello-grpc/pkg/api/v1"
)

type HelloStreamService struct {
}

func (s *HelloStreamService) BiDirectStream(stream v1.HelloStreamService_BiDirectStreamServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(in)
		if err := stream.Send(&v1.StreamDataResponse{
			Api: "v1",
			StreamData: &v1.StreamData{
				Id:    in.StreamData.Id + 1,
				Value: in.StreamData.Value + 1,
			},
		}); err != nil {
			return err
		}
	}

	return nil
}

func (s *HelloStreamService) ClientStream(stream v1.HelloStreamService_ClientStreamServer) error {

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Println(in)
	}

	if err := stream.SendMsg(&v1.StreamDataResponse{
		Api: "v1",
		StreamData: &v1.StreamData{
			Id:    1,
			Value: 2,
		},
	}); err != nil {
		return err
	}

	return nil
}

func (s *HelloStreamService) ServerStream(req *v1.StreamDataRequest, stream v1.HelloStreamService_ServerStreamServer) error {

	product, ok := products[req.StreamData.Id]
	if !ok {
		return errors.Errorf("product %d not found", req.StreamData.Id)
	}

	for i := 0; i < 100; i++ {
		if err := stream.Send(&v1.StreamDataResponse{
			Api: "v1",
			StreamData: &v1.StreamData{
				Id:    req.StreamData.Id,
				Value: product.Price.Number,
			},
		}); err != nil {
			return err
		}
	}
	return nil
}
