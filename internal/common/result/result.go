package result

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	errx2 "github.com/zouchangfu/go-zero-element-admin/internal/common/errx"
	"net/http"
)

func HttpResult(w http.ResponseWriter, r *http.Request, resp interface{}, err error) {
	if err != nil {
		errCode := errx2.ServerCommonError
		errMsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)
		if e, ok := causeErr.(*errx2.CodeError); ok {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		}
		logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
		return
	}

	result := Success(resp)
	httpx.WriteJson(w, http.StatusOK, result)
	return
}

type Result struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(data interface{}) *Result {
	return &Result{200, "OK", data}
}

func Error(errCode uint32, errMsg string) *Result {
	return &Result{Code: errCode, Msg: errMsg}
}
