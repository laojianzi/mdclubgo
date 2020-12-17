package register

import (
	"errors"
	"fmt"
	"html"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/db"
	"github.com/laojianzi/mdclubgo/internal/avatar"
	"github.com/laojianzi/mdclubgo/internal/database"
	"github.com/laojianzi/mdclubgo/internal/email/validator"
	"github.com/laojianzi/mdclubgo/internal/exception"
	"github.com/laojianzi/mdclubgo/internal/present"
	"github.com/laojianzi/mdclubgo/log"
)

// Form user register request json body
type Form struct {
	Email     string `json:"email"`
	EmailCode string `json:"email_code"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

// Validate for register form data
func (f Form) Validate() *exception.MDClubGoError {
	mdclubgoErr := exception.ErrFieldVerifyFailed
	if f.Email == "" {
		mdclubgoErr = mdclubgoErr.AddErrors("email", "email不能为空")
	} else {
		user := &database.User{Email: f.Email}
		exists, err := database.Has(db.Instance(), user, func(g *gorm.DB) *gorm.DB {
			return g.Where(user)
		})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exists = false
			err = nil
		}

		if err != nil {
			log.Error(fmt.Errorf("has user email: %w", err).Error())
			return exception.ErrInternalServerError
		}

		if exists {
			mdclubgoErr = mdclubgoErr.AddErrors("email", "该邮箱已被注册")
		}
	}

	// skip check email code if not installed
	if conf.Installed {
		if f.EmailCode == "" {
			mdclubgoErr = mdclubgoErr.AddErrors("email_code", "email_code不能为空")
		} else {
			ok, err := validator.CheckCode(f.Email, f.EmailCode)
			if errors.Is(err, validator.ErrEmailVerifyExpired) {
				return exception.ErrAPIMethodNotAllowed
			}

			if err != nil {
				log.Error(fmt.Errorf("register valid: %w", err).Error())
				return exception.ErrInternalServerError
			}

			if !ok {
				mdclubgoErr = mdclubgoErr.AddErrors("email_code", "邮箱验证码错误")
			}
		}
	}

	if f.Username == "" {
		mdclubgoErr = mdclubgoErr.AddErrors("username", "username不能为空")
	} else {
		user := &database.User{Username: f.Username}
		exists, err := database.Has(db.Instance(), user, func(g *gorm.DB) *gorm.DB {
			return g.Where(user)
		})
		if errors.Is(err, gorm.ErrRecordNotFound) {
			exists = false
			err = nil
		}

		if err != nil {
			log.Error(fmt.Errorf("has user username: %w", err).Error())
			return exception.ErrInternalServerError
		}

		if exists {
			mdclubgoErr = mdclubgoErr.AddErrors("username", "该用户名已被注册")
		}
	}

	if f.Password == "" {
		mdclubgoErr = mdclubgoErr.AddErrors("password", "password不能为空")
	}

	if len(mdclubgoErr.Errors) > 0 {
		return mdclubgoErr
	}

	return nil
}

// Present user register response json body
type Present struct {
	ctx echo.Context
	*database.User
	Avatars       map[string]string `json:"avatar"`
	Covers        map[string]string `json:"cover"`
	Relationships struct {
		IsMe        bool `json:"is_me"`
		IsFollowing bool `json:"is_following"`
		IsFollowed  bool `json:"is_followed"`
	} `json:"relationships"`
}

// NewPresent return a *register.Present
func NewPresent(ctx echo.Context, user *database.User) *Present {
	return &Present{ctx: ctx, User: user}
}

// Format return a response format data
func (p Present) Format() interface{} {
	u := p.User
	if u.ID > 0 {
		if u.Avatar != "" {
			p.Avatars = avatar.BrandUrls(p.ctx, new(avatar.UserAvatar), u.ID, u.Avatar)
		}

		p.Covers = avatar.BrandUrls(p.ctx, new(avatar.UserCover), u.ID, u.Cover)
	}

	if u.Headline != "" {
		u.Headline = html.EscapeString(u.Headline)
	}

	if u.Bio != "" {
		u.Bio = html.EscapeString(u.Bio)
	}

	if u.Blog != "" {
		u.Blog = html.EscapeString(u.Blog)
	}

	if u.Company != "" {
		u.Company = html.EscapeString(u.Company)
	}

	if u.Location != "" {
		u.Location = html.EscapeString(u.Location)
	}

	return present.Build(p)
}
