package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	myutils "src/tools.app/utils"
)

func TestSendMail(t *testing.T) {
	t.Skip("skip TestSendMail.")
	t.Log("Case01: test send an email.")
	entry := &myutils.MailEntry{
		ServerPwd: "*******",
		MailTo:    []string{"zhengjin@4paradigm.com"},
		Subject:   "Go Mail Test",
		Body:      "This is a go mail test.",
	}
	if err := myutils.SendMail(entry); err != nil {
		t.Fatal(err)
	}
}

func TestSendMailWithAttachFiles(t *testing.T) {
	t.Skip("skip TestSendMailWithAttachFiles.")
	t.Log("Case01: test send an email with attached files.")
	baseDir := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
	files := []string{
		filepath.Join(baseDir, "test.out"),
		filepath.Join(baseDir, "test.json"),
	}

	entry := &myutils.MailEntry{
		ServerPwd:   "*******",
		MailTo:      []string{"zhengjin@4paradigm.com"},
		Subject:     "Go Mail Test",
		Body:        "This is a go mail test with attached files.",
		AttachFiles: files,
		IsArchive:   false,
	}

	if err := myutils.SendMail(entry); err != nil {
		t.Fatal(err)
	}
}

func TestSendMailWithAttachArchive(t *testing.T) {
	t.Skip("skip TestSendMailWithAttachArchive.")
	t.Log("Case01: test send an email with attached archive file.")
	baseDir := filepath.Join(os.Getenv("HOME"), "Downloads/tmp_files")
	files := []string{
		filepath.Join(baseDir, "test_log.txt"),
		filepath.Join(baseDir, "logs"),
	}

	entry := &myutils.MailEntry{
		ServerPwd:   "*******",
		MailTo:      []string{"zhengjin@4paradigm.com"},
		Subject:     "Go Mail Test",
		Body:        "This is a go mail test with attached archive file.",
		AttachFiles: files,
		IsArchive:   true,
	}

	if err := myutils.SendMail(entry); err != nil {
		t.Fatal(err)
	}
}
