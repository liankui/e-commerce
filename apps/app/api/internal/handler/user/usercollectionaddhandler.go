package user

import (
	"net/http"

	"github.com/liankui/e-commerce/apps/app/api/internal/logic/user"
	"github.com/liankui/e-commerce/apps/app/api/internal/svc"
	"github.com/liankui/e-commerce/apps/app/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserCollectionAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCollectionAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewUserCollectionAddLogic(r.Context(), svcCtx)
		resp, err := l.UserCollectionAdd(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
