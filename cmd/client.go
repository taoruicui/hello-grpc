package cmd

import (
	"context"
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
				Unit:       v1.Product_Price_USD,
				Number:     2.01,
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
