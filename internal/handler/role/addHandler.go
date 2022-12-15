package role

import (
	"github.com/zouchangfu/go-zero-element-admin/internal/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zouchangfu/go-zero-element-admin/internal/logic/role"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RoleAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := role.NewAddLogic(r.Context(), svcCtx)
		err := l.Add(&req)
		result.HttpResult(w, r, nil, err)
	}
}
