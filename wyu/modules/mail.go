package modules

import (
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"runtime"
)

var Htmls []string

type MailConfigs struct {
	Username string
	Password string
	Host string
	Port int
}

type MailSendParams struct {
	From string
	To string
	Subject string
	Body string
}

type Mail struct {
	d *gomail.Dialer
	cfg *MailConfigs
}

func NewMail(cfg *MailConfigs) (m *Mail) {
	m = &Mail{cfg:cfg}
	m.loading()
	return
}

func (mail *Mail) loading() *Mail {
	if mail.cfg != nil {
		mail.d = gomail.NewDialer(mail.cfg.Host, mail.cfg.Port, mail.cfg.Username, mail.cfg.Password)
	}

	return mail
}

func (mail *Mail) Send(params *MailSendParams) (err error) {
	if mail.d == nil {
		_, file, line, _ := runtime.Caller(1)
		err = errors.New(fmt.Sprintf("Plz Initialized Mail! » %v » %v", file, line))
		return
	}

	if params == nil {
		_, file, line, _ := runtime.Caller(1)
		err = errors.New(fmt.Sprintf("Plz add Parameters to Send Mail! » %v » %v", file, line))
		return
	}

	msg := gomail.NewMessage()
	msg.SetHeader("From", params.From)
	msg.SetHeader("To", params.To)
	msg.SetHeader("Subject", params.Subject)
	msg.SetBody("text/html", params.Body)

	return mail.d.DialAndSend(msg)
}

func (mail *Mail) Tpls() (err error) {
	return
}


