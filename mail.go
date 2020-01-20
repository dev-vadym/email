package email

// IMail store function in email package
type IMail interface {
	Send(from, recipient, subject, msg string) error
}

type Mail struct {
	mailClient IMail
	from       string
	recipient  string
	subject    string
	msg        string
}

func InitMail(m IMail, from, recipient, subject, msg string) *Mail {
	return &Mail{
		mailClient: m,
		from:       from,
		recipient:  recipient,
		subject:    subject,
		msg:        msg,
	}
}

func (m *Mail) Sent() {
	m.mailClient.Send(m.from, m.recipient, m.subject, m.msg)
}
