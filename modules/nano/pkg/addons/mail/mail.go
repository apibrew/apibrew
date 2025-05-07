package mail

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/dop251/goja"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/semaphore"
	gomail "gopkg.in/mail.v2"
)

type mailObject struct {
	config    EmailConfig
	semaphore *semaphore.Weighted
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
	FromName    string `json:"fromName"`
	To          string `json:"to"`
	ToName      string `json:"toName"`
	Subject     string `json:"subject"`
	Body        string `json:"body"`
	ContentType string `json:"contentType"`
}

func mailFn(config EmailConfig) *mailObject {
	obj := &mailObject{config: config, semaphore: semaphore.NewWeighted(10)}

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

func (m *mailObject) Send(message EmailMessage) bool {
	mail := gomail.NewMessage()

	if message.FromName != "" {
		mail.SetHeader("From", fmt.Sprintf("%s <%s>", message.FromName, message.From))
	} else {
		mail.SetHeader("From", message.From)
	}
	if message.ToName != "" {
		mail.SetHeader("To", fmt.Sprintf("%s <%s>", message.ToName, message.To))
	} else {
		mail.SetHeader("To", message.To)
	}
	mail.SetHeader("Subject", message.Subject)
	mail.SetBody(message.ContentType, message.Body)

	d := gomail.NewDialer(m.config.Host, m.config.Port, m.config.Username, m.config.Password)

	if m.config.EnableTls {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	}

	err := d.DialAndSend(mail)

	if err != nil {
		log.Error(err)
		return false
	}
	return true
}

func (m *mailObject) SendParallel(message EmailMessage) bool {
	log.Println("Begin sending email to: " + message.To)
	if err := m.semaphore.Acquire(context.TODO(), 1); err != nil {
		log.Println(err)
	}

	go func() {
		defer func() {
			m.semaphore.Release(1)
			log.Println("End sending email to: " + message.To)
		}()
		mail := gomail.NewMessage()

		if message.FromName != "" {
			mail.SetHeader("From", fmt.Sprintf("%s <%s>", message.FromName, message.From))
		} else {
			mail.SetHeader("From", message.From)
		}
		if message.ToName != "" {
			mail.SetHeader("To", fmt.Sprintf("%s <%s>", message.ToName, message.To))
		} else {
			mail.SetHeader("To", message.To)
		}
		mail.SetHeader("Subject", message.Subject)
		mail.SetBody(message.ContentType, message.Body)

		d := gomail.NewDialer(m.config.Host, m.config.Port, m.config.Username, m.config.Password)

		if m.config.EnableTls {
			d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		}

		err := d.DialAndSend(mail)

		if err != nil {
			log.Error(err)
		}
	}()
	return true
}

func Register(vm *goja.Runtime) error {
	return vm.Set("mailer", mailFn)
}
