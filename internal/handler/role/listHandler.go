package role

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zouchangfu/go-zero-element-admin/internal/common/result"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
	"net/http"

	"github.com/zouchangfu/go-zero-element-admin/internal/logic/role"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RolePageReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List(&req)
		result.HttpResult(w, r, resp, err)
	}
}
