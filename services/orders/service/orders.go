package service

import ("context"
orderproto "gomicro/services/common/genproto/orders")


type Orderservice struct{
	//store
}
var ordersDB = make([]*orderproto.Order,0)

func NewOrderService() *Orderservice{
	return &Orderservice{}
}
//rpc CreateOder(CreateOrderRequest) returns (CreateOrderResponse);

func (o *Orderservice) CreateOder(ctx context.Context,order *orderproto.Order)error{
	ordersDB = append(ordersDB,order)
	return nil

}
func (o *Orderservice)CreateOrderResponse(){
	
}