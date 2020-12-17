package utils

import (
	"fmt"
	"net/smtp"

	configs "github.com/yerzhan-karatay/survey-webapp-backend/config"
)

var emailAuth smtp.Auth

// SendEmailSMTP is for sending email
func SendEmailSMTP(to []string, emailBody string) (bool, error) {
	config := configs.Get()
	smtpServer := config.SMTP.Server
	emailFrom := config.SMTP.Email
	smtpPassword := config.SMTP.Password
	smtpPort := config.SMTP.Port

	emailAuth = smtp.PlainAuth("", emailFrom, smtpPassword, smtpServer)

	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + "Your response from Survey World!\n"
	msg := []byte(subject + mime + "\n" + emailBody)
	addr := fmt.Sprintf("%s:%s", smtpServer, smtpPort)

	if err := smtp.SendMail(addr, emailAuth, emailFrom, to, msg); err != nil {
		return false, err
	}
	return true, nil
}
