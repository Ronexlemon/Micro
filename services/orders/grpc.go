package orders

import (
	handler "gomicro/services/orders/handler/order"
	"gomicro/services/orders/service"
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
	orderService := service.NewOrderService()
    handler.NewGrpcOrdersService(grpcServer,orderService)

	return grpcServer.Serve(lis)
}