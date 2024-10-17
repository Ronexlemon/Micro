package main

import "gomicro/services/orders"

func main(){
	grpcServer := orders.NewGRPCServer(":9090")
	grpcServer.Run()
}