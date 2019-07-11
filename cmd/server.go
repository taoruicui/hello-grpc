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

		// register rpc services
		api.RegisterHelloServiceServer(s.GrpcServer, &svc.ServiceServer{})
		api.RegisterHelloStreamServiceServer(s.GrpcServer, &svc.HelloStreamService{})

		// register gateway for http endpoints
		gwMux := runtime.NewServeMux()
		err := api.RegisterHelloServiceHandlerFromEndpoint(s.Context, gwMux, fmt.Sprintf("localhost:%d", s.Port), []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = api.RegisterHelloStreamServiceHandlerFromEndpoint(s.Context, gwMux, fmt.Sprintf("localhost:%d", s.Port), []grpc.DialOption{grpc.WithInsecure()})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		s.Register(gwMux)

		// start server
		err = s.Run()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
