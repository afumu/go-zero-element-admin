package user

import (
	"github.com/zouchangfu/go-zero-element-admin/common/result"
	"github.com/zouchangfu/go-zero-element-admin/internal/logic/user"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func QueryByIdHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FormParamId
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewQueryByIdLogic(r.Context(), svcCtx)
		resp, err := l.QueryById(&req)
		result.HttpResult(w, r, resp, err)
	}
}
