package utils_test

import (
	"testing"

	myutils "tools.app/utils"
)

func TestSendMail(t *testing.T) {
	t.Log("Case01: test send an email.")
	entry := &myutils.MailEntry{
		MailTo:   []string{"zhengjin@4paradigm.com"},
		Subject:  "go mail test",
		Body:     "this is a go mail test.",
		ConnPass: "*******",
	}
	if err := myutils.SendMail(entry); err != nil {
		t.Fatal(err)
	}
}
