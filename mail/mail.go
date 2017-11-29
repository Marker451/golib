package mail

import (
	"gopkg.in/gomail.v2"
	"log"
	"net/smtp"
)

const (
	HOST        = "smtp.163.com"
	SERVER_ADDR = "smtp.163.com:465"
	USER        = "youxileiting@163.com"  //发送邮件的邮箱
	PASSWORD    = ""                      //发送邮件邮箱的密码
	TO          = "zhongjizhizun@126.com" //默认收件
)

func SendEmailWithGomail(to string, subject string, msg string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", USER)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(HOST, 465, USER, PASSWORD)
	d.SSL = true
	d.Auth = unencryptedAuth{
		smtp.PlainAuth("", USER, PASSWORD, HOST),
	}

	if err := d.DialAndSend(m); err != nil {
		log.Println("send gomail err ", err)
		return err
	}
	return nil
}

type unencryptedAuth struct {
	smtp.Auth
}

func (a unencryptedAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	s := *server
	s.TLS = true
	return a.Auth.Start(&s)
}
