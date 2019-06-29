package gotests_test

import (
	"bytes"
	"fmt"
	"testing"

	"demo.tests/gotests"
)

// cmd: go test -v src/demo.tests/gotests/echo_test.go
func TestEcho(t *testing.T) {
	var tests = []struct {
		newline bool
		sep     string
		args    []string
		want    string
	}{
		{true, "", []string{}, "\n"},
		{false, "", []string{}, ""},
		{true, "\t", []string{"one", "two", "three"}, "one\ttwo\tthree\n"},
		{true, ",", []string{"a", "b", "c"}, "a,b,c\n"},
		{false, ":", []string{"1", "2", "3"}, "1:2:3"},
		{true, ",", []string{"a", "b", "c"}, "a b c\n"}, // NOTE: failed expectation!
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v, %q, %q)", test.newline, test.sep, test.args)
		gotests.Out = new(bytes.Buffer)

		if err := gotests.Echo(test.newline, test.sep, test.args); err != nil {
			t.Errorf("%s failed: %v", descr, err)
			continue
		}
		got := gotests.Out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
