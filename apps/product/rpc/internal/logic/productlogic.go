package logic

import (
	"context"
	"fmt"
	"github.com/liankui/e-commerce/apps/product/rpc/model"

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
	v, err, _ := l.svcCtx.SingleGroup.Do(fmt.Sprintf("product:%d", in.ProductId), func() (interface{}, error) {
		return l.svcCtx.ProductModel.FindOne(l.ctx, in.ProductId)
	})
	if err != nil {
		return nil, err
	}
	p := v.(*model.Product)
	return &product.ProductItem{
		ProductId: p.Id,
		Name:      p.Name,
		Stock:     p.Stock,
	}, nil
}
