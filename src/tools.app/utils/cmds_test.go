package utils_test

import (
	"testing"

	myutils "tools.app/utils"
)

func TestRunShellCmd(t *testing.T) {
	cmd := "df -h"
	t.Logf("Case01: test run shell command: %s\n", cmd)
	output1, err := myutils.RunShellCmd(cmd)
	if err != nil {
		t.Fatalf("run command (%s) failed: %v\n", cmd, err)
	}
	if len(output1) == 0 {
		t.Error("failed: command is empty!")
	}
	t.Logf("command %s output: %s\n", cmd, output1)

	t.Logf("Case02: test run shell command with output bufferred: %s\n", cmd)
	output2, err := myutils.RunShellCmdBuf(cmd)
	if err != nil {
		t.Fatalf("run command (%s) failed: %v\n", cmd, err)
	}
	if len(output2) == 0 {
		t.Error("failed: command is empty!")
	}
	t.Logf("command %s output: %s\n", cmd, output2)
}

func TestRunShellCmds(t *testing.T) {
	cmds := []string{"hostname", "df -h"}
	output, err := myutils.RunShellCmds(cmds)
	if err != nil {
		t.Fatalf("run command (%s) failed: %v\n", cmds, err)
	}
	if len(output) == 0 {
		t.Error("failed: command is empty!")
	}
	t.Logf("command %v output: %s\n", cmds, output)
}
