package handler

import (
	"context"
	orderproto "gomicro/services/common/genproto/orders"
	types "gomicro/services/orders/type"

	"google.golang.org/grpc"
)


type OrdersGrpcHandler struct{
	ordersservice  types.OrderService
	
	orderproto.UnimplementedOrderserviceServer
}

func NewGrpcOrdersService(grpc *grpc.Server,ordersService  types.OrderService){
	gRPCHandler := &OrdersGrpcHandler{ordersservice: ordersService}

	//register the orders service
	orderproto.RegisterOrderserviceServer(grpc, gRPCHandler)
	
	
}

func(h *OrdersGrpcHandler) CreateOder(ctx context.Context,request *orderproto.CreateOrderRequest)(*orderproto.CreateOrderResponse,error){
	order := &orderproto.Order{
		OrderID: 43,
		CustomerId: 2,
		ProductID: 1,
		Quantity: 2,
	}
	err:=h.ordersservice.CreateOder(ctx,order)
	if err!=nil{
		return nil,err
	}
	 //customerID, productId ,quantity := request.GetCustomerID(),request.GetProductId(),request.GetQuantity()

	 return &orderproto.CreateOrderResponse{Status: "Success"},nil
    
     

}