package types

import (
	"context"
	orderproto "gomicro/services/common/genproto/orders"
)

type OrderService interface{
	CreateOder(context.Context,*orderproto.Order) error
}
