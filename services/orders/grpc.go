package orders

import (
	"log"
	"net"

	"google.golang.org/grpc"
)


type gRPCServer struct{
	addr string
}

func NewGRPCServer(addr string)*gRPCServer{
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error{
	lis,err := net.Listen( "tcp", s.addr)
	if err !=nil{
		log.Fatal("Failed to listen",err)
	}
	grpcServer := grpc.NewServer()

	return grpcServer.Serve(lis)
}