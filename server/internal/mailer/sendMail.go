package mailer

import (
	"an-overengineered-app/internal/config"
	"an-overengineered-app/internal/logger"
	"context"
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

func getSubjectTitle(mailType string) string {
	switch mailType {
	case "auth":
		return "An overengineered app - Verify your email to continue"
	default:
		return "An overengineered app"
	}
}

func SendMail(ctx context.Context, receiver string, content []byte, mailType string) error {

	logger.PrintInfo(ctx, "Sending mail", map[string]string{
		"mailType": mailType,
		"receiver": receiver,
	})
	emailConfig := config.EmailConfig

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	mailer := gomail.NewDialer(emailConfig.SMTPServer, emailConfig.Port, "", "")
	mailer.TLSConfig = tlsConfig

	mail := gomail.NewMessage()
	mail.SetHeader("From", emailConfig.From)
	mail.SetHeader("To", receiver)
	mail.SetHeader("Subject", getSubjectTitle(mailType))
	mail.SetBody("text/html", string(content))

	if err := mailer.DialAndSend(mail); err != nil {
		logger.PrintErrorWithStack(ctx, "Failed to send email", err)
		return err
	}

	return nil

}
