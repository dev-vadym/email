# email
````go
import "github.com/golang-common-packages/email"

// This is for both SendGrid and GoMail
mailClient := email.NewMailClient(email.SENDGRID, &email.MailConfig{
    URL:       "",
    Port:      "",
    Username:  "",
    Password:  "",
    SecretKey: "",
}),

if err := mailClient.Send("", "", "", ""); err != nil {
    log.Printf("Can not send: %s", err.Error())
}
````