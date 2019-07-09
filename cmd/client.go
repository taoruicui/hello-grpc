package cmd

import (
	"context"
	"io"
	"log"

	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/spf13/cobra"
	"github.com/taoruicui/hello-grpc/pkg/api/v1"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(clientCmd)
	clientCmd.AddCommand(streamCmd)
}

var clientCmd = &cobra.Command{
	Use: "client",
	Run: func(cmd *cobra.Command, args []string) {
		RunClient()
	},
}

func RunClient() {

	conn, err := grpc.Dial("localhost:8887", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// create request
	createReq := v1.ProductRequest{
		Api: "v1",
		Product: &v1.Product{
			Id: 1,
			Price: &v1.Product_Price{
				Name:       "product 1",
				Unit:       v1.Product_Price_CNY,
				Number:     12.32,
				UploadedAt: ptypes.TimestampNow(),
			},
		},
	}

	res, err := c.Create(ctx, &createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	log.Println(res)

	// read request
	readReq := v1.ProductRequest{
		Api: "v1",
		Product: &v1.Product{
			Id: 1,
		},
	}

	resHello, err := c.Read(ctx, &readReq)
	if err != nil {
		log.Fatalf("Read failed: %v", err)
	}

	log.Println(proto.MarshalTextString(resHello))

	// update request
	updateReq := v1.ProductRequest{
		Api: "v1",
		Product: &v1.Product{
			Id: 1,
			Price: &v1.Product_Price{
				Unit:      v1.Product_Price_USD,
				Number:    2.01,
				UpdatedAt: ptypes.TimestampNow(),
			},
		},
	}

	res, err = c.Update(ctx, &updateReq)
	if err != nil {
		log.Fatalf("update failed: %v", err)
	}

	log.Println(res)

	// create request
	createReq = v1.ProductRequest{
		Api: "v1",
		Product: &v1.Product{
			Id: 2,
			Price: &v1.Product_Price{
				Name:       "product 2",
				Unit:       v1.Product_Price_YEN,
				Number:     100,
				UploadedAt: ptypes.TimestampNow(),
			},
		},
	}

	res, err = c.Create(ctx, &createReq)
	if err != nil {
		log.Fatalf("Create failed: %v", err)
	}

	log.Println(res)

	// delete request
	deleteReq := v1.ProductRequest{
		Api: "v1",
		Product: &v1.Product{
			Id: 2,
		},
	}

	res, err = c.Delete(ctx, &deleteReq)
	if err != nil {
		log.Fatalf("delete failed: %v", err)
	}

	log.Println(res)

}

var streamCmd = &cobra.Command{
	Use: "stream",
	Run: func(cmd *cobra.Command, args []string) {
		mode := args[0]
		if mode == "bi" {
			RunBiSteam()
		} else if mode == "client" {
			RunClientStream()
		} else {
			RunServerStream()
		}
	},
}

func RunServerStream() {
	conn, err := grpc.Dial("localhost:8887", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewHelloStreamServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	serverStream, err := c.ServerStream(ctx, &v1.StreamDataRequest{
		Api: "v1",
		StreamData: &v1.StreamData{
			Id: 1,
		},
	})
	if err != nil {
		log.Fatalf("Send failed: %v", err)
	}

	for {
		res, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res)
	}

}

func RunClientStream() {

	conn, err := grpc.Dial("localhost:8887", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewHelloStreamServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientStream, err := c.ClientStream(ctx)
	if err != nil {
		log.Fatalf("%v.clientStream(_) = _, %v", clientStream, err)
	}
	for i := 0; i < 100; i++ {
		if err := clientStream.Send(&v1.StreamDataRequest{
			Api: "v1",
			StreamData: &v1.StreamData{
				Id:    int64(i),
				Value: float64(i),
			},
		}); err != nil {
			log.Fatalf("%v.Send = %v", clientStream, err)
		}
	}

	reply, err := clientStream.CloseAndRecv()
	if err == nil {
		log.Printf("clientStream summary: %v", reply)
	} else {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", clientStream, err, nil)
	}
}

func RunBiSteam() {
	conn, err := grpc.Dial("localhost:8887", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := v1.NewHelloStreamServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	biStream, err := c.BiDirectStream(ctx)
	if err != nil {
		log.Fatalf("%v.biStream(_) = _, %v", biStream, err)
	}

	go func(stream v1.HelloStreamService_BiDirectStreamClient) {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Println(res)
		}
	}(biStream)

	for i := 0; i < 100; i++ {
		if err := biStream.Send(&v1.StreamDataRequest{
			Api: "v1",
			StreamData: &v1.StreamData{
				Id:    int64(i),
				Value: float64(i),
			},
		}); err != nil {
			log.Fatalf("%v.Send = %v", biStream, err)
		}
	}
}
