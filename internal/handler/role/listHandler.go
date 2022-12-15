package role

import (
	"github.com/zouchangfu/go-zero-element-admin/internal/common/result"
	"net/http"

	"github.com/zouchangfu/go-zero-element-admin/internal/logic/role"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := role.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		result.HttpResult(w, r, resp, err)
	}
}
