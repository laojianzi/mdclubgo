package register

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/laojianzi/mdclubgo/email"
	emailtemplate "github.com/laojianzi/mdclubgo/internal/email/template"
	"github.com/laojianzi/mdclubgo/log"
)

// Send register email content
func Send(to string, appName, username string) error {
	if to == "" {
		return fmt.Errorf("register email send to can't empty")
	}

	if appName == "" {
		return fmt.Errorf("register email send app name can't empty")
	}

	if username == "" {
		return fmt.Errorf("register email send username can't empty")
	}

	return email.Send([]string{to}, Content(to, appName, username))
}

// Content get register content
func Content(to, appName, username string) string {
	data := struct {
		To       string
		AppName  string
		Username string
	}{
		To:       to,
		AppName:  appName,
		Username: username,
	}

	tmplContent := string(emailtemplate.MustAsset("conf/email/register.tmpl"))
	tpl, err := template.New("register").Parse(tmplContent)
	if err != nil {
		log.Panic(fmt.Errorf("new register template: %w", err).Error())
	}

	buf := bytes.NewBufferString("")
	err = tpl.Execute(buf, data)
	if err != nil {
		log.Panic(fmt.Errorf("register template execute: %w", err).Error())
	}

	return buf.String()
}
