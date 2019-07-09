package cmd

import (
	"fmt"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	api "github.com/taoruicui/hello-grpc/pkg/api/v1"
	"github.com/taoruicui/hello-grpc/pkg/server"
	svc "github.com/taoruicui/hello-grpc/pkg/service/v1"
	"google.golang.org/grpc"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use: "server",
	Run: func(cmd *cobra.Command, args []string) {

		s := server.DefaultServer()
		api.RegisterHelloServiceServer(s.GrpcServer, &svc.ServiceServer{})
		api.RegisterHelloStreamServiceServer(s.GrpcServer, &svc.HelloStreamService{})

		mux := runtime.NewServeMux()
		err := api.RegisterHelloServiceHandlerFromEndpoint(s.Context, mux, fmt.Sprintf("localhost:%d", s.Port), []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = api.RegisterHelloStreamServiceHandlerFromEndpoint(s.Context, mux, fmt.Sprintf("localhost:%d", s.Port), []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		s.HttpServer.Handler = mux

		err = s.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
