package gotests

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Out : global var, and write test results to Out
var Out io.Writer = os.Stdout

// Echo : join and display input arguments
func Echo(newline bool, sep string, args []string) error {
	fmt.Fprint(Out, strings.Join(args, sep))
	if newline {
		fmt.Fprintln(Out)
	}
	return nil
}
