package main

import (
	gpb "github.com/DQFSN/blog/proto/grpc"
	"github.com/DQFSN/blog/server/rpcimpl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v",err)
	}

	s := grpc.NewServer()
	gpb.RegisterAuthServer(s, &rpcimpl.Auth{})
	gpb.RegisterPublishServer(s, &rpcimpl.BlogServer{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}