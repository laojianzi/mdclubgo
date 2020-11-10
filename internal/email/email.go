package email

import (
	"fmt"
	"regexp"

	"github.com/laojianzi/mdclubgo/internal/captcha"
	"github.com/laojianzi/mdclubgo/internal/exception"
	"github.com/laojianzi/mdclubgo/internal/present"
	"github.com/laojianzi/mdclubgo/log"
)

// Present result struct for email
type Present struct{}

// Format return a response format data
func (p Present) Format() interface{} {
	return present.Build(nil)
}

// Form receive form for email
type Form struct {
	Email string `json:"email"`
	captcha.Form
}

// Validate for email form data
func (f Form) Validate() *exception.MDClubGoError {
	mdclubgoErr := exception.NewMDClubGoError(exception.StatusBadRequest, exception.MessageBadRequest)
	emailMatched, err := regexp.MatchString(`^[a-zA-Z0-9.!#$%&'*+/=?^_`+
		"`"+`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`, f.Email)
	if err != nil {
		log.Error("email regexp match error: %s", err.Error())
		mdclubgoErr = mdclubgoErr.AddErrors("email", fmt.Sprintf("Regexp 错误: %s", err.Error()))
	}

	if err == nil && !emailMatched {
		mdclubgoErr = mdclubgoErr.AddErrors("email", "不能为空")
	}

	if len(mdclubgoErr.Errors) > 0 {
		return mdclubgoErr
	}

	return nil
}
