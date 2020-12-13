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

// define http code for common error
const (
	StatusFieldVerifyFailed = iota + 200001
	StatusSendEmailFailed
	StatusEmailVerifyExpired
	StatusImageUploadFailed
	StatusImageNotFound
	StatusVoteTypeError
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

// define http message for common error
const (
	MessageFieldVerifyFailed  = "字段验证失败"
	MessageSendEmailFailed    = "邮件发送失败"
	MessageEmailVerifyExpired = "邮件验证码已失效"
	MessageImageUploadFailed  = "图片上传失败"
	MessageImageNotFound      = "指定图片不存在"
	MessageVoteTypeError      = "投票类型只能是 up、down 中的一个"
)

// MDClubGoError error for mdclubgo
type MDClubGoError struct {
	Code    MDClubGoErrorCode    `json:"code"`
	Message MDClubGoErrorMessage `json:"message"`
	Errors  map[string]string    `json:"errors,omitempty"` // map[error_field]error_message pls uses `AddErrors` add
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
		Errors:  make(map[string]string),
	}
}

// AddErrors add to `Errors`
func (err *MDClubGoError) AddErrors(key, msg string) *MDClubGoError {
	err.Errors[key] = msg
	return err
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

var (
	ErrFieldVerifyFailed  = NewMDClubGoError(StatusFieldVerifyFailed, MessageFieldVerifyFailed)
	ErrSendEmailFailed    = NewMDClubGoError(StatusSendEmailFailed, MessageSendEmailFailed)
	ErrEmailVerifyExpired = NewMDClubGoError(StatusEmailVerifyExpired, MessageEmailVerifyExpired)
	ErrImageUploadFailed  = NewMDClubGoError(StatusImageUploadFailed, MessageImageUploadFailed)
	ErrImageNotFound      = NewMDClubGoError(StatusImageNotFound, MessageImageNotFound)
	ErrVoteTypeError      = NewMDClubGoError(StatusVoteTypeError, MessageVoteTypeError)
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
