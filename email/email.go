package email

import (
	"fmt"
	"strings"

	"github.com/laojianzi/mdclubgo/conf"
	"github.com/laojianzi/mdclubgo/email/smtp"
	"github.com/laojianzi/mdclubgo/log"
)

var instance Mailer

// Mailer mail interface
type Mailer interface {
	Send(to []string, msg string) error
}

// Init for Mailer
func Init() {
	typ := strings.ToLower(strings.TrimSpace(conf.Email.Type))

	switch typ {
	case SMTP:
		instance = smtp.NewSMTP(smtp.WithHost(conf.Email.Host), smtp.WithFrom(conf.Email.From),
			smtp.WithUsername(conf.Email.Username), smtp.WithPassword(conf.Email.Password))
	case POP3:
		// TODO: add POP3
	case IMAP:
		// TODO: add IMAP
	default:
		log.Fatal(fmt.Errorf("email type unrecognized dialect: %s", typ).Error())
	}
}

// Send mail
func Send(to []string, msg string) error {
	return instance.Send(to, msg)
}

// SetTestMailer set instance for testing
func SetTestMailer(mailer Mailer) {
	instance = mailer
}
