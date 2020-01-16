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
func NewMailClient(mailClientType int, config *MailConfig) IMailClient {
	switch mailClientType {
	case GOMAIL:
		return NewGoEmailClient(config.URL, config.Port, config.Username, config.Password)
	case SENDGRID:
		return NewSendGridClient(config.SecretKey)
	}

	return nil
}
