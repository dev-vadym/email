package email

// IMailClient store function in email package
type IMailClient interface {
	Send(from, recipient, subject, msg string) error
}

const (
	GOMAIL = iota
	SENDGRID
)

// NewMailClient function for Factory Pattern
func NewMailClient(mailClientType int, url, portNumber, username, password, keyService string) IMailClient {
	switch mailClientType {
	case GOMAIL:
		return NewGoEmailClient(url, portNumber, username, password)
	case SENDGRID:
		return NewSendGridClient(keyService)
	}

	return nil
}
