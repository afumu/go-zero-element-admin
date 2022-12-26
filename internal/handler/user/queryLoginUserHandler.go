package user

import (
	"github.com/zouchangfu/go-zero-element-admin/internal/common/result"
	"net/http"

	"github.com/zouchangfu/go-zero-element-admin/internal/logic/user"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
)

func QueryLoginUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewQueryLoginUserLogic(r.Context(), svcCtx)
		resp, err := l.QueryLoginUser()
		result.HttpResult(w, r, resp, err)
	}
}
