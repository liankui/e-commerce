package svc

import (
	"github.com/liankui/e-commerce/apps/app/api/internal/config"
	"github.com/liankui/e-commerce/apps/order/rpc/order"
	"github.com/liankui/e-commerce/apps/product/rpc/product"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
	}
}
