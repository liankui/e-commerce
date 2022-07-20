package logic

import (
	"context"

	"github.com/liankui/e-commerce/apps/product/rpc/internal/svc"
	"github.com/liankui/e-commerce/apps/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type OperationProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOperationProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OperationProductsLogic {
	return &OperationProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *OperationProductsLogic) OperationProducts(in *product.OperationProductsRequest) (*product.OperationProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.OperationProductsResponse{}, nil
}
