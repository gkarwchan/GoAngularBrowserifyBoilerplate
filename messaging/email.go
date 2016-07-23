package messaging

import (
	"github.com/gkarwchan/GoAngularBrowserifyBoilerplate/app"
	"gopkg.in/gomail.v1"
)

var (
	mailer *gomail.Mailer
)

func init() {
	mailer = gomail.NewMailer(app.SmtpHost, app.SmtpUsername, app.SmtpPassword, app.SmtpPort)
}

// SendEmail will send email from template
func SendEmail(body, recipient, subject string) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "system@myacne.expert")
	msg.SetHeader("To", recipient)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)
	return mailer.Send(msg)
}
