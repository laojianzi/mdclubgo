package captcha

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/jpeg"
	"net/http"
	"strings"

	"github.com/dchest/captcha"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/laojianzi/mdclubgo/cache"
	"github.com/laojianzi/mdclubgo/internal/exception"
	"github.com/laojianzi/mdclubgo/internal/present"
	"github.com/laojianzi/mdclubgo/limiter"
	"github.com/laojianzi/mdclubgo/log"
)

// LifeTime captcha validity period
const LifeTime = 3600

// CacheKeyPrefix captcha cache key prefix
const CacheKeyPrefix = "captcha_"

// Present result struct for captcha
type Present struct {
	CaptchaImage string `json:"captcha_image"`
	CaptchaToken string `json:"captcha_token"`
}

// Format return a response format data
func (p Present) Format() interface{} {
	return present.Build(p)
}

// Form receive form for captcha
type Form struct {
	CaptchaToken string `json:"captcha_token"`
	CaptchaCode  string `json:"captcha_code"`
}

// CacheKey get captcha cache key
func CacheKey(key string) string {
	return fmt.Sprintf("%s%s", CacheKeyPrefix, key)
}

// Generate return a *Captcha
// set width = 100 if width value is 0
// set height = 36 if height value is 0
func Generate(width, height int) (*Present, error) {
	if width == 0 {
		width = 100
	}

	if height == 0 {
		height = 36
	}

	code := captcha.RandomDigits(5)
	image := captcha.NewImage(captcha.NewLen(5), code, width, height)
	token := middleware.DefaultRequestIDConfig.Generator()
	err := cache.Set(CacheKey(token), string(code), LifeTime)
	if err != nil {
		return nil, fmt.Errorf("captcha cache set: %w", err)
	}

	imageContent := bytes.NewBufferString("")
	err = jpeg.Encode(imageContent, image, &jpeg.Options{Quality: 90})
	if err != nil {
		return nil, fmt.Errorf("captcha write image buffer: %w", err)
	}

	// base64 jepg
	jpegImage := fmt.Sprintf("data:image/jpeg;base64,%s", base64.StdEncoding.EncodeToString(imageContent.Bytes()))
	return &Present{CaptchaImage: jpegImage, CaptchaToken: token}, nil
}

// Check validity for captcha token and code
func Check(token, code string) bool {
	if token == "" || code == "" {
		return false
	}

	currentCode := cache.Get(CacheKey(token), "")
	if currentCode == "" {
		return false
	}

	if err := cache.Delete(CacheKey(token)); err != nil {
		log.Error(fmt.Errorf("captcha check: %w", err).Error())
	}

	return strings.EqualFold(currentCode, code)
}

// NextTimeNeed does the next request require a captcha?
func NextTimeNeed(ctx echo.Context, id, action string, maxCount, period int) bool {
	remaining := limiter.ActLimit(ctx, id, action, maxCount, period)
	needCaptcha := remaining <= 1

	form := new(Form)
	err := ctx.Bind(form)
	if err != nil {
		log.Error(fmt.Errorf("next time need bind: %w", err).Error())
	}

	if remaining <= 0 && !Check(form.CaptchaToken, form.CaptchaCode) {
		err := exception.ErrFieldVerifyFailed.AddErrors("captcha_code", "图形验证码错误")
		ctx.Error(echo.NewHTTPError(http.StatusForbidden).SetInternal(err))
		return false
	}

	return needCaptcha
}
