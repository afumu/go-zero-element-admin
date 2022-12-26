package dict

import (
	"github.com/zouchangfu/go-zero-element-admin/internal/common/result"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zouchangfu/go-zero-element-admin/internal/logic/dict"
	"github.com/zouchangfu/go-zero-element-admin/internal/svc"
	"github.com/zouchangfu/go-zero-element-admin/internal/types"
)

func EditHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DictEditReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := dict.NewEditLogic(r.Context(), svcCtx)
		err := l.Edit(&req)
		result.HttpResult(w, r, nil, err)
	}
}
