package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/logic/user"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/types"
)

func RemoveByIdsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FormParamIds
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRemoveByIdsLogic(r.Context(), svcCtx)
		resp, err := l.RemoveByIds(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
