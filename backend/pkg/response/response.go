package response

import (
	"backend/pkg/code"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type SuccessBean struct {
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(data interface{}) *SuccessBean {
	return &SuccessBean{code.OK, "OK", data}
}

type ErrorBean struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func Error(errCode uint32, errMsg string) *ErrorBean {
	return &ErrorBean{errCode, errMsg}
}

func HttpResult(w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		res := Success(resp)
		httpx.WriteJson(w, http.StatusOK, res)
		return
	}
	errCode := code.ServerInternalError
	errMsg := code.UnknownError
	causeErr := errors.Cause(err)
	var e *code.Error
	if errors.As(causeErr, &e) {
		errCode = e.GetErrCode()
		errMsg = e.GetErrMsg()
	}
	if errCode != code.IgnoredError {
		logx.Errorf("【API-ERR】%+v ", err)
	}
	httpx.WriteJson(w, http.StatusBadRequest, Error(errCode, errMsg))
}

func ParamErrorResult(w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s: %s", code.ErrorMessage(code.RequestParamError), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(code.RequestParamError, errMsg))
}
