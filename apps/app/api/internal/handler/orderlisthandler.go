package handler

import (
	"net/http"

	"github.com/liankui/e-commerce/apps/app/api/internal/logic"
	"github.com/liankui/e-commerce/apps/app/api/internal/svc"
	"github.com/liankui/e-commerce/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func OrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewOrderListLogic(r.Context(), svcCtx)
		resp, err := l.OrderList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
