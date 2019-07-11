package utils

// dependency:
// git submodule add "https://github.com/go-gomail/gomail.git" src/tools.app/vendor/gopkg.in/gomail.v2
import (
	"fmt"

	gomail "gopkg.in/gomail.v2"
)

// MailEntry entry for send an mail.
type MailEntry struct {
	MailTo   []string
	Subject  string
	Body     string
	ConnPass string
}

// SendMail sends a mail by smtp.
func SendMail(entry *MailEntry) error {
	const (
		connUser = "zivieira@163.com"
		connHost = "smtp.163.com"
		connPort = 25
	)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("ZJ Test<%s>", connUser)) // add a tag
	m.SetHeader("To", entry.MailTo...)
	m.SetHeader("Subject", entry.Subject)
	m.SetBody("text/html", entry.Body)

	d := gomail.NewDialer(connHost, connPort, connUser, entry.ConnPass)
	err := d.DialAndSend(m)
	return err
}
