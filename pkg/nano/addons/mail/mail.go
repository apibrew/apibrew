package mail

import (
	"crypto/tls"
	"github.com/dop251/goja"
)
import gomail "gopkg.in/mail.v2"

type mailObject struct {
	config EmailConfig
}

type EmailConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	EnableAuth bool   `json:"enableAuth"`
	EnableTls  bool   `json:"enableTls"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type EmailMessage struct {
	From        string `json:"from"`
	To          string `json:"to"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	ContentType string `json:"contentType"`
}

func mailFn(config EmailConfig) *mailObject {
	obj := &mailObject{config: config}

	if obj.config.Host == "" {
		panic("host is required")
	}

	if obj.config.Port == 0 {
		panic("port is required")
	}

	if obj.config.EnableAuth && obj.config.Username == "" {
		panic("username is required")
	}

	if obj.config.EnableAuth && obj.config.Password == "" {
		panic("password is required")
	}

	return obj
}

func (m *mailObject) Send(message EmailMessage) {
	mail := gomail.NewMessage()

	mail.SetHeader("From", message.From)
	mail.SetHeader("To", message.To)
	mail.SetHeader("Subject", message.Subject)
	mail.SetBody(message.ContentType, message.Body)

	d := gomail.NewDialer(m.config.Host, m.config.Port, m.config.Username, m.config.Password)

	if m.config.EnableTls {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	err := d.DialAndSend(mail)

	if err != nil {
		panic(err)
	}
}

func Register(vm *goja.Runtime) error {
	return vm.Set("mailer", mailFn)
}
