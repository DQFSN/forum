package main

import (
	"github.com/DQFSN/forum/config"
	gpb "github.com/DQFSN/forum/proto/grpc"
	"github.com/DQFSN/forum/server/rpcimpl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var (
	port string
)
func init() {
	gRPCConfig := config.Get().GRPC
	port = gRPCConfig.Port
}


func main() {

	lis, err := net.Listen("tcp", ":"+port)
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