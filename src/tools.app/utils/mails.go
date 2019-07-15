package utils

// dependency:
// git submodule add "https://github.com/go-gomail/gomail.git" src/tools.app/vendor/gopkg.in/gomail.v2
import (
	"fmt"
	"os"

	gomail "gopkg.in/gomail.v2"
)

// MailEntry entry for send an mail.
type MailEntry struct {
	Meta        string   `json:"meta"`
	ServerPwd   string   `json:"pwd"`
	MailTo      []string `json:"receivers"`
	Subject     string   `json:"subject"`
	Body        string   `json:"body"`
	AttachFiles []string `json:"attachments"`
	IsArchive   bool     `json:"archive"`
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

	if entry.IsArchive {
		f, err := createArchiveFile(entry.AttachFiles)
		if err != nil {
			return err
		}
		m.Attach(f)
	} else {
		for _, f := range entry.AttachFiles {
			m.Attach(f)
		}
	}

	d := gomail.NewDialer(connHost, connPort, connUser, entry.ServerPwd)
	return d.DialAndSend(m)
}

func createArchiveFile(paths []string) (string, error) {
	var files []*os.File
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return "", err
		}
		files = append(files, f)
	}

	outputArchive := fmt.Sprintf("/tmp/archive_%s.tar.gz", GetCurrentDateTime())
	err := CompressGzipFile(files, outputArchive)
	if err != nil {
		return "", err
	}
	return outputArchive, err
}
