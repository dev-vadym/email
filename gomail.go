package email

import (
	"crypto/tls"
	"log"
	"strconv"

	"gopkg.in/gomail.v2"
)

// GoMailClient manage all GoMail action
type GoMailClient struct {
	dialer     *gomail.Dialer
	session    gomail.SendCloser
	url        string
	portNumber string
	username   string
	password   string
}

// NewGoEmailClient function return GoMail client based on singleton pattern
func NewGoEmailClient(url, portNumber, username, password string) IMailClient {
	currentSession := &GoMailClient{nil, nil, "", "", "", ""}

	dialer, err := getDialer(url, portNumber, username, password)
	if err != nil {
		log.Println("Error when try to make strconv port from config: ", err)
		panic(err)
	}

	session, err := dialer.Dial()
	if err != nil {
		log.Println("Error when try to dial to mail server: ", err)
		panic(err)
	}
	log.Println("Connected to Mail Server")

	currentSession.dialer = dialer
	currentSession.session = session
	currentSession.url = url
	currentSession.portNumber = portNumber
	currentSession.username = username
	currentSession.password = password

	return currentSession
}

// Send function sent mail based on argument provide
func (e *GoMailClient) Send(from, to, subject, message string) (err error) {
	msg := gomail.NewMessage()

	msg.SetHeader("From", from)
	msg.SetAddressHeader("To", to, "")
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", message)

	if e.session != nil {
		if err = e.session.Send(from, []string{to}, msg); err != nil {
			// close current session
			e.session.Close()

			// generate new client
			e.dialer, _ = getDialer(e.url, e.portNumber, e.username, e.password)

			// reconnect smtp server
			if newSession, err := e.dialer.Dial(); err == nil {
				e.session = newSession
				// resend email
				return e.session.Send(from, []string{to}, msg)
			}
			return err
		}
		return nil
	}

	return e.dialer.DialAndSend(msg)
}

// getDialer private function return a new Dialer
func getDialer(url, portNumber, username, password string) (client *gomail.Dialer, err error) {
	port, err := strconv.Atoi(portNumber)
	if err != nil {
		return nil, err
	}

	client = gomail.NewDialer(url, port, username, password)
	client.TLSConfig = &tls.Config{InsecureSkipVerify: true, ServerName: url}
	client.LocalName = url

	return client, nil
}
