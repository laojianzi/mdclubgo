package exception

import (
	"fmt"
	"net/http"
)

// MDClubGoErrorCode define error code for mdclubgo
type MDClubGoErrorCode uint

// define http code for system error
const (
	StatusInternalServerError = iota + 100000
	StatusServiceUnavailable
	StatusIPTooManyRequests
	StatusUserTooManyRequests
	StatusAPINotFound
	StatusAPIMethodNotAllowed
	StatusBadRequest
	StatusSysInstallFailed
)

// MDClubGoErrorMessage define error message for mdclubgo
type MDClubGoErrorMessage string

// define http message for system error
const (
	MessageInternalServerError = "服务器错误"
	MessageServiceUnavailable  = "系统维护中"
	MessageIPTooManyRequests   = "IP 请求超过上限"
	MessageUserTooManyRequests = "用户请求超过上限"
	MessageAPINotFound         = "接口不存在"
	MessageAPIMethodNotAllowed = "该接口不支持此 HTTP METHOD"
	MessageBadRequest          = "请求参数的 json 格式错误"
	MessageSysInstallFailed    = "系统安装失败"
)

// MDClubGoError error for mdclubgo
type MDClubGoError struct {
	Code    MDClubGoErrorCode    `json:"code"`
	Message MDClubGoErrorMessage `json:"message"`
}

// Error format MDClubGoError to string
func (err *MDClubGoError) Error() string {
	return fmt.Sprintf("code=%d, message=%s", err.Code, err.Message)
}

// NewMDClubGoError return a *MDClubGoError
func NewMDClubGoError(code MDClubGoErrorCode, msg MDClubGoErrorMessage) *MDClubGoError {
	return &MDClubGoError{
		Code:    code,
		Message: msg,
	}
}

// define http error for system
var (
	ErrInternalServerError = NewMDClubGoError(StatusInternalServerError, MessageInternalServerError)
	ErrServiceUnavailable  = NewMDClubGoError(StatusServiceUnavailable, MessageServiceUnavailable)
	ErrIPTooManyRequests   = NewMDClubGoError(StatusIPTooManyRequests, MessageIPTooManyRequests)
	ErrUserTooManyRequests = NewMDClubGoError(StatusUserTooManyRequests, MessageUserTooManyRequests)
	ErrAPINotFound         = NewMDClubGoError(StatusAPINotFound, MessageAPINotFound)
	ErrAPIMethodNotAllowed = NewMDClubGoError(StatusAPIMethodNotAllowed, MessageAPIMethodNotAllowed)
	ErrBadRequest          = NewMDClubGoError(StatusBadRequest, MessageBadRequest)
	ErrSysInstallFailed    = NewMDClubGoError(StatusSysInstallFailed, MessageSysInstallFailed)
)

var httpCodeToMDClubGoError = map[int]*MDClubGoError{
	http.StatusInternalServerError: ErrInternalServerError,
	http.StatusBadGateway:          ErrServiceUnavailable,
	http.StatusServiceUnavailable:  ErrServiceUnavailable,
	http.StatusTooManyRequests:     ErrIPTooManyRequests, // default is ip too many requests
	http.StatusNotFound:            ErrAPINotFound,
	http.StatusMethodNotAllowed:    ErrAPIMethodNotAllowed,
	http.StatusBadRequest:          ErrBadRequest,
}

// HTTPCodeToMDClubGoError get a *MDClubGoError by http code
func HTTPCodeToMDClubGoError(code int) *MDClubGoError {
	return httpCodeToMDClubGoError[code]
}