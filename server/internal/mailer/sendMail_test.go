package mailer

import (
	"an-overengineered-app/internal/config"
	"context"
	"fmt"
	"testing"

	smtpmock "github.com/mocktools/go-smtp-mock"
	"github.com/stretchr/testify/assert"
)

func TestGetSubjectTitle(t *testing.T) {
	t.Run("Should return title for auth if mailType is given as auth", func(t *testing.T) {
		title := getSubjectTitle("auth")

		assert.Equal(t, "An overengineered app - Verify your email to continue", title)
	})

	t.Run("Should return default title if no mailType provided", func(t *testing.T) {
		title := getSubjectTitle("")

		assert.Equal(t, "An overengineered app", title)
	})
}

func startMockServer(host string, port int) *smtpmock.Server {
	mockServer := smtpmock.New(smtpmock.ConfigurationAttr{
		HostAddress: host,
		PortNumber:  port,
	})
	go mockServer.Start()
	return mockServer
}

func configureEmailConfig(server string, port int, from string) {
	config.EmailConfig.SMTPServer = server
	config.EmailConfig.Port = port
	config.EmailConfig.From = from
}

func setMailData() (string, []byte) {
	receiver := "user@test.com"
	mailContent := []byte("Test mail content")
	return receiver, mailContent
}

func TestSendMail(t *testing.T) {

	t.Run("Should send email on provided host", func(t *testing.T) {
		mockServer := startMockServer("127.0.0.1", 1026)

		defer mockServer.Stop()

		receiver, mailContent := setMailData()

		configureEmailConfig("127.0.0.1", 1026, receiver)

		err := SendMail(context.TODO(), receiver, mailContent, "auth")
		assert.NoError(t, err)

		assert.Equal(t, len(mockServer.Messages()), 1)
	})

	t.Run("Should throw error if email fails to sent", func(t *testing.T) {
		mockServer := startMockServer("", 0)

		defer mockServer.Stop()

		receiver, mailContent := setMailData()

		configureEmailConfig("", 0, "")

		err := SendMail(context.TODO(), receiver, mailContent, "auth")
		fmt.Println(err)

		assert.Error(t, err)
	})

}
