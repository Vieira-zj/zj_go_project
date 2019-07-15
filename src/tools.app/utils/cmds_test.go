package utils_test

import (
	"strings"
	"testing"

	myutils "tools.app/utils"
)

func TestRunShellCmd(t *testing.T) {
	t.Skip("skip TestRunShellCmd.")
	// TODO: output is empty for command: "java -version"
	// cmd := "java -version"
	cmd := "go version"
	t.Logf("Case01: test run shell command: %s\n", cmd)
	output1, err := myutils.RunShellCmd(cmd)
	if err != nil {
		// Logf();FailNow()
		t.Fatalf("Failed run command (%s): %v\n", cmd, err)
	}
	if len(output1) == 0 {
		// Log();Fail()
		t.Error("Failed: command output is empty!")
	}
	t.Logf("command %s output: %s\n", cmd, output1)

	t.Logf("Case02: test run shell command with output bufferred: %s\n", cmd)
	output2, err := myutils.RunShellCmdBuf(cmd)
	if err != nil {
		t.Fatalf("Failed run command (%s): %v\n", cmd, err)
	}
	if len(output2) == 0 {
		t.Error("Failed: command output is empty!")
	}
	t.Logf("command (%s) output: %s\n", cmd, output2)
}

func TestRunShellCmds(t *testing.T) {
	t.Skip("skip TestRunShellCmds.")
	cmds := []string{"hostname", "go version", "ls /tmp"}
	output, err := myutils.RunShellCmds(cmds)
	if err != nil {
		t.Fatalf("run commands (%s) failed: %v\n", strings.Join(cmds, ","), err)
	}
	if len(output) == 0 {
		t.Error("failed: command output is empty!")
	}
	t.Logf("commands (%s) output: %s\n", strings.Join(cmds, ","), output)
}
