package user

import (
	"github.com/zouchangfu/go-zero-element-admin/common/result"
	"net/http"

	"github.com/zouchangfu/go-zero-element-admin/service/internal/logic/user"
	"github.com/zouchangfu/go-zero-element-admin/service/internal/svc"
)

func ListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewListLogic(r.Context(), svcCtx)
		resp, err := l.List()
		result.HttpResult(w, r, resp, err)
	}
}
