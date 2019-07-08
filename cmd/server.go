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

		mux := runtime.NewServeMux()
		err := api.RegisterHelloServiceHandlerFromEndpoint(s.Context, mux, fmt.Sprintf("localhost:%d", s.Port), []grpc.DialOption{grpc.WithInsecure()})
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

//func start(port string) error {
//
//	v1API := v1Svc.NewServiceServer()
//
//	// register service
//	server := grpc.NewServer()
//	v1.RegisterHelloServiceServer(server, v1API)
//
//	// graceful shutdown
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt)
//	go func() {
//		for range c {
//			// sig is a ^C, handle it
//			log.Println("shutting down gRPC server...")
//
//			server.GracefulStop()
//
//			<-ctx.Done()
//		}
//	}()
//
//	// start gRPC server
//	log.Println("starting gRPC server...")
//	return server.Serve(listen)
//
//	//mux := http.NewServeMux()
//	//if err := addRoutes(mux, server); err != nil {
//	//	log.Fatalf("adding routes: %v", err)
//	//}
//	//httpServer := &http.Server{
//	//	Handler: mux,
//	//}
//
//}
//
//func startHttpServer(port string) error {
//	ctx := context.Background()
//	ctx, cancel := context.WithCancel(ctx)
//	defer cancel()
//
//	mux := runtime.NewServeMux()
//	opts := []grpc.DialOption{grpc.WithInsecure()}
//	err := v1.RegisterHelloServiceHandlerFromEndpoint(ctx, mux, "localhost:8888", opts)
//	if err != nil {
//		return err
//	}
//
//	c := make(chan os.Signal, 1)
//	signal.Notify(c, os.Interrupt)
//	go func() {
//		for range c {
//			// sig is a ^C, handle it
//			log.Println("shutting down http server...")
//
//			<-ctx.Done()
//		}
//	}()
//	return http.ListenAndServe(":8887", mux)
//}
