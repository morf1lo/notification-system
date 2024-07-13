package service

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/morf1lo/notification-system/internal/worker/model"
)

type MailerService struct {
	from string
	pass string
	host string
	port string
}

func NewMailerService() *MailerService {
	return &MailerService{
		from: os.Getenv("EMAIL"),
		pass: os.Getenv("EMAIL_PASS"),
		host: os.Getenv("SMTP_HOST"),
		port: os.Getenv("SMTP_PORT"),
	}
}

func (s *MailerService) Send(to string, message *model.Article) error {
	subject := message.Title
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
		</head>
		<body>
			<h2>%s</h2>
			<h3>Author - %s</h3>
			<p>%s</p>
		</body>
		</html>
	`, message.Title, message.Author, message.Body)

	msg := []byte("Subject: " + subject + "\r\n" +
	"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
	"\r\n" + body)

	auth := smtp.PlainAuth("", s.from, s.pass, s.host)

	err := smtp.SendMail(s.host + ":" + s.port, auth, s.from, []string{to}, msg)
	return err
}
