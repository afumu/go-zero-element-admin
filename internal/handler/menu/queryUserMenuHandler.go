package menu

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zouchangfu/go-zero-element-admin/internal/logic/menu"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
)

func QueryUserMenuHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := menu.NewQueryUserMenuLogic(r.Context(), svcCtx)
		resp, err := l.QueryUserMenu()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
