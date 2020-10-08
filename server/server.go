package main

import (
	pb "blog/api"
	"blog/server/rpcimpl"
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
	pb.RegisterAuthServer(s, &rpcimpl.Auth{})
	pb.RegisterPublishServer(s, &rpcimpl.BlogServer{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v",err)
	}
}