package code

import (
	"fmt"
	"strings"
)

// OK 成功返回
const OK uint32 = 200

// 前3位代表业务,后三位代表具体功能
const (
	IgnoredError uint32 = 100000 + iota
	ServerInternalError
	RequestParamError
)

var message map[uint32]string

func init() {
	message = map[uint32]string{
		IgnoredError:        "ignored error",
		ServerInternalError: "internal error",
		RequestParamError:   "invalid param",
	}
}

const UnknownError = "unknown error"

func ErrorMessage(code uint32) string {
	if m, ok := message[code]; ok {
		return m
	}
	return UnknownError
}

type Error struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 返回给前端的错误码
func (e *Error) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 返回给前端显示端错误信息
func (e *Error) GetErrMsg() string {
	return e.errMsg
}

func (e *Error) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewError(errCode uint32, errMsg string) *Error {
	return &Error{errCode: errCode, errMsg: errMsg}
}

func NewInternalError(errMsg ...string) *Error {
	if len(errMsg) > 0 {
		return NewError(ServerInternalError, strings.Join(errMsg, " "))
	}
	return NewError(ServerInternalError, ErrorMessage(ServerInternalError))
}

func NewIgnoredError(errMsg string) *Error {
	return NewError(IgnoredError, errMsg)
}

func NewInvalidParamError() *Error {
	return NewError(RequestParamError, ErrorMessage(RequestParamError))
}
