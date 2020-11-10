package smtp

import (
	"net"
	"net/smtp"
)

// Mail smtp adapter
type Mail struct {
	username string
	password string
	host     string
	from     string
}

// Option option mail field
type Option func(*Mail)

// WithUsername set username
func WithUsername(username string) Option {
	return func(m *Mail) {
		m.username = username
	}
}

// WithPassword set password
func WithPassword(password string) Option {
	return func(m *Mail) {
		m.password = password
	}
}

// WithHost set host
func WithHost(host string) Option {
	return func(m *Mail) {
		m.host = host
	}
}

// WithFrom set from
func WithFrom(from string) Option {
	return func(m *Mail) {
		m.from = from
	}
}

// NewSMTP return a opened *Mail
func NewSMTP(opt ...Option) *Mail {
	m := new(Mail)
	for _, o := range opt {
		o(m)
	}

	return m
}

// Send use smtp send email
// to := []string{"recipient@example.net"}
// msg := []byte("To: recipient@example.net\r\n" +
//    "Subject: discount Gophers!\r\n" +
//    "\r\n" +
//    "This is the email body.\r\n")
// err := smtp.SendMail("mail.example.com:25", auth, "sender@example.org", to, msg)
// if err != nil {
//    log.Fatal(err)
// }
func (m *Mail) Send(to []string, msg string) error {
	host, port, err := net.SplitHostPort(m.host)
	if err != nil {
		return err
	}

	auth := smtp.PlainAuth("", m.username, m.password, host)
	return smtp.SendMail(net.JoinHostPort(host, port), auth, m.from, to, []byte(msg))
}
