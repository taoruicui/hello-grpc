package server

import (
	"net"
	"os"

	"context"

	"fmt"

	"os/signal"
	"syscall"

	"net/http"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	GrpcServer *grpc.Server
	HttpServer *http.Server
	Context    context.Context
	Port       int
}

func NewServer() *Server {
	// TODO new server from config
	return &Server{
		GrpcServer: grpc.NewServer(),
		Context:    context.Background(),
		Port:       8887,
	}
}

func DefaultServer() *Server {
	return &Server{
		GrpcServer: grpc.NewServer(),
		HttpServer: &http.Server{
			Addr: "localhost:8888",
		},
		Context: context.Background(),
		Port:    8887,
	}
}

func (s *Server) Init() {

}

func (s *Server) Run() error {

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		return err
	}

	// start gRPC server
	go func() {
		log.Infof("gRPC server listen to localhost:%d", s.Port)
		if err := s.GrpcServer.Serve(listen); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	// start http server
	go func() {
		log.Infof("http server listen to %s", s.HttpServer.Addr)
		if err := s.HttpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start http server: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	select {
	// wait on kill signal
	case <-ch:
		log.Info("kill signal received")
		s.Stop()
		// wait on context cancel
	case <-s.Context.Done():
		log.Info("context cancel received")
		s.Stop()
	}

	return nil
}

func (s *Server) Stop() {

	log.Info("Server is shutting down...")

	s.GrpcServer.GracefulStop()
	s.HttpServer.Shutdown(s.Context)

	log.Info("Server is down")

}
