package validator

import (
	"bytes"
	"crypto/md5" // nolint:gosec
	"fmt"
	"strconv"
	"strings"
	"text/template"

	"github.com/laojianzi/mdclubgo/cache"
	"github.com/laojianzi/mdclubgo/email"
	emailtemplate "github.com/laojianzi/mdclubgo/internal/email/template"
	"github.com/laojianzi/mdclubgo/internal/stringutil"
	"github.com/laojianzi/mdclubgo/log"
)

const (
	LifeTime       = 3600 * 3      // 邮件有效期 3 小时
	MaxTimes       = 20            // 每个邮件验证码最多验证次数
	CacheKeyPrefix = "email_code_" // email validator cache key prefix
)

// CacheKey get captcha cache key
func CacheKey(key string) string {
	return fmt.Sprintf("%s%s", CacheKeyPrefix, fmt.Sprintf("%x", md5.Sum([]byte(key)))) // nolint:gosec
}

// GenerateCode generate validator code
func GenerateCode(email string) string {
	code, times := CodeInfo(email)
	if code != "" && times > 0 {
		return code
	}

	code = stringutil.Random(6, nil)
	err := cache.Set(CacheKey(email), fmt.Sprintf("%s-0", code), LifeTime)
	if err != nil {
		log.Panic(fmt.Errorf("cache set validator code info: %w", err).Error())
	}

	return code
}

// CodeInfo return validator code and times
func CodeInfo(email string) (string, int) {
	codeInfo := cache.Get(CacheKey(email), "")
	if codeInfo == "" {
		return "", 0
	}

	info := strings.Split(codeInfo, "-")
	code := info[0]
	times, err := strconv.Atoi(info[1])
	if err != nil {
		log.Panic("code info times '%s' not int", info[1])
	}

	return code, times
}

// Send validator email content
func Send(to string, appName, code string) error {
	if to == "" {
		return fmt.Errorf("validator send to can't empty")
	}

	if appName == "" {
		return fmt.Errorf("validator send app name can't empty")
	}

	if code == "" {
		return fmt.Errorf("validator send code can't empty")
	}

	return email.Send([]string{to}, Content(to, appName, code))
}

// Content get validator content
func Content(to, appName, code string) string {
	data := struct {
		To      string
		AppName string
		Code    string
	}{
		To:      to,
		AppName: appName,
		Code:    code,
	}

	globalTemp := string(emailtemplate.MustAsset("conf/email/validator.tmpl"))
	tpl, err := template.New("validator").Parse(globalTemp)
	if err != nil {
		log.Panic(fmt.Errorf("new validator template: %w", err).Error())
	}

	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, data)
	if err != nil {
		log.Panic(fmt.Errorf("validator template execute: %w", err).Error())
	}

	return buf.String()
}
