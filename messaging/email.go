package messaging

import (
	"github.com/gkar68/GoAngularBrowserifyBoilerplate/settings"
	"gopkg.in/gomail.v1"
)

var (
	mailer *gomail.Mailer
)


func init() {
	mailer = gomail.NewMailer(settings.SmtpHost, settings.SmtpUsername, settings.SmtpPassword, settings.SmtpPort)
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
