package logic

import (
	"context"

	"github.com/liankui/e-commerce/apps/product/rpc/internal/svc"
	"github.com/liankui/e-commerce/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductLogic {
	return &ProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductLogic) Product(in *product.ProductItemRequest) (*product.ProductItem, error) {
	// todo: add your logic here and delete this line

	return &product.ProductItem{}, nil
}