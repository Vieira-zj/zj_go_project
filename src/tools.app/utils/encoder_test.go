package utils_test

import (
	"fmt"
	"strings"
	"testing"

	myutils "src/tools.app/utils"
)

func TestGetBase64Text(t *testing.T) {
	const testText = "Hello World!"
	t.Log("Case01: test get base64 md5 text")
	retText := myutils.GetBase64Text([]byte(testText))
	if len(retText) == 0 {
		t.Error("Failed: returned md5 hex text is empty!")
	}

	md5Cmd := fmt.Sprintf("echo -n '%s' | base64", testText)
	if expectText, err := myutils.RunShellCmd(md5Cmd); err == nil {
		if strings.Trim(expectText, "\n") != retText {
			t.Errorf("Failed: base64 text not matched.")
		}
	}
	t.Logf("base64 text: %s", retText)
}

func TestGetMd5HexText(t *testing.T) {
	const testText = "Hello World!"
	t.Log("Case01: test get md5 hex text")
	retText := myutils.GetMd5HexText(testText)
	if len(retText) == 0 {
		t.Error("Failed: returned md5 hex text is empty!")
	}

	md5Cmd := fmt.Sprintf("echo -n '%s' | md5sum | awk '{print $1}'", testText)
	if expectText, err := myutils.RunShellCmd(md5Cmd); err == nil {
		if strings.Trim(expectText, "\n") != retText {
			t.Errorf("Failed: md5 text not matched.")
		}
	}
	t.Logf("md5 text: %s", retText)
}
