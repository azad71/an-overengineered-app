package mailer

import (
	"an-overengineered-app/internal/config"
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func getSubjectTitle(mailType string) string {
	switch mailType {
	case "auth":
		return "An overengineered social app - Verify your email to continue"
	default:
		return "An overengineered social app"
	}
}

func SendMail(receiver []string, content []byte, mailType string) error {
	emailConfig := config.EmailConfig

	newMail := &email.Email{
		From:    emailConfig.From,
		To:      receiver,
		HTML:    content,
		Subject: getSubjectTitle(mailType),
	}

	mailUrl := fmt.Sprintf("%s:%d", emailConfig.SMTPServer, emailConfig.Port)

	// fmt.Printf("Constructed mail url: %s\n", mailUrl)

	err := newMail.Send(mailUrl, smtp.PlainAuth("", "Admin", "", emailConfig.SMTPServer))

	if err != nil {
		fmt.Printf("Failed to send email. Error: %v", err)
		return err
	}

	return nil

}
