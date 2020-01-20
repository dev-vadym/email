package email

// IMail interface for strategy pattern
type IMail interface {
	Send(from, recipient, subject, msg string) error
}

// Mail for context embeds the strategy
type Mail struct {
	mailClient IMail
	from       string
	recipient  string
	subject    string
	msg        string
}

func InitMailClient(client IMail, from, recipient, subject, msg string) *Mail {
	return &Mail{
		mailClient: client,
		from:       from,
		recipient:  recipient,
		subject:    subject,
		msg:        msg,
	}
}

func (m *Mail) SetMailClient(client IMail) {
	m.mailClient = client
}

func (m *Mail) SetMailContext(from, recipient, subject, msg string) {
	m.from = from
	m.recipient = recipient
	m.subject = subject
	m.msg = msg
}

func (m *Mail) Sent() {
	m.mailClient.Send(m.from, m.recipient, m.subject, m.msg)
}
