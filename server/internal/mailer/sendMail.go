package mailer

import (
	"an-overengineered-app/internal/config"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/gomail.v2"
)

func getSubjectTitle(mailType string) string {
	switch mailType {
	case "auth":
		return "An overengineered social app - Verify your email to continue"
	default:
		return "An overengineered social app"
	}
}

func SendMail(receiver string, content []byte, mailType string) error {
	emailConfig := config.EmailConfig

	certPath, _ := os.Getwd()
	certPath = filepath.Join(certPath, "certs", "smtp", "cert.pem")

	cert, err := os.ReadFile(certPath)

	if err != nil {
		log.Fatalf("Failed to read smtp cert. Error: %v", err)
		return err
	}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(cert)
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	mailer := gomail.NewDialer(emailConfig.SMTPServer, emailConfig.Port, "", "")
	mailer.TLSConfig = tlsConfig

	mail := gomail.NewMessage()
	mail.SetHeader("From", emailConfig.From)
	mail.SetHeader("To", receiver)
	mail.SetHeader("Subject", getSubjectTitle(mailType))
	mail.SetBody("text/html", string(content))

	if err := mailer.DialAndSend(mail); err != nil {
		fmt.Printf("Failed to send email. Error: %v", err)
		return err
	}

	return nil

}
