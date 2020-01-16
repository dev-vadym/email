package email

import (
	"log"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGridClient manage all email action
type SendGridClient struct {
	client *sendgrid.Client
}

// NewSendGridClient function return sendgrid client based on singleton pattern
func NewSendGridClient(keyService string) IMailClient {
	currentSession := &SendGridClient{nil}

	client := sendgrid.NewSendClient(keyService)
	currentSession.client = client
	log.Println("Connected to SendGrid Server")

	return currentSession
}

// Send function sent mail based on argument provide
func (s *SendGridClient) Send(from, recipient, subject, msg string) error {
	From := mail.NewEmail("Go-Common-Package", from)
	To := mail.NewEmail("Reviewer", recipient)
	plainTextContent := msg
	htmlContent := "<strong>" + msg + "</strong>"
	message := mail.NewSingleEmail(From, subject, To, plainTextContent, htmlContent)
	_, err := s.client.Send(message)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
