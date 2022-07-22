package svc

import (
	"github.com/liankui/e-commerce/apps/order/rpc/internal/config"
	"github.com/liankui/e-commerce/apps/order/rpc/model"
	"github.com/liankui/e-commerce/apps/product/rpc/product"
	"github.com/liankui/e-commerce/apps/user/rpc/user"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	OrderModel     model.OrdersModel
	OrderitemModel model.OrderitemModel
	ShippingModel  model.ShippingModel
	UserRpc        user.User
	ProductRpc     product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:         c,
		OrderModel:     model.NewOrdersModel(conn, c.CacheRedis),
		OrderitemModel: model.NewOrderitemModel(conn, c.CacheRedis),
		ShippingModel:  model.NewShippingModel(conn, c.CacheRedis),
		UserRpc:        user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc:     product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
