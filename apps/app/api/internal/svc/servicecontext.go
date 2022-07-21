package svc

import (
	"github.com/liankui/e-commerce/apps/app/api/internal/config"
	"github.com/liankui/e-commerce/apps/order/rpc/order"
	"github.com/liankui/e-commerce/apps/product/rpc/product"
	"github.com/liankui/e-commerce/apps/reply/rpc/reply"
	"github.com/liankui/e-commerce/apps/user/rpc/user"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	OrderRPC   order.Order
	ProductRPC product.Product
	ReplyRPC   reply.Reply
	UserRPC    user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRPC:   order.NewOrder(zrpc.MustNewClient(c.OrderRPC)),
		ProductRPC: product.NewProduct(zrpc.MustNewClient(c.ProductRPC)),
		ReplyRPC:   reply.NewReply(zrpc.MustNewClient(c.ReplyRPC)),
		UserRPC:    user.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}
