package utils

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

// RunShellCmd runs shell command.
func RunShellCmd(cmd string) (string, error) {
	cmmd := exec.Command("/bin/bash", "-c", cmd)
	stdout, err := cmmd.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := cmmd.Start(); err != nil {
		return "", err
	}
	b, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	if err := cmmd.Wait(); err != nil {
		return "", err
	}
	return string(b), nil
}

// RunShellCmdBuf runs shell command and stores output bytes in buffer.
func RunShellCmdBuf(cmd string) (string, error) {
	cmmd := exec.Command("/bin/bash", "-c", cmd)
	stdout, err := cmmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := cmmd.Start(); err != nil {
		return "", err
	}

	var lines []string
	outputBuf := bufio.NewReader(stdout)
	for {
		b, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				return "", err
			}
			break
		}
		lines = append(lines, string(b))
	}

	if err := cmmd.Wait(); err != nil {
		return "", err
	}
	return strings.Join(lines, "\n"), nil
}

// RunShellCmds run multiple shell commands in shell client.
func RunShellCmds(cmds []string) (string, error) {
	in := bytes.NewBuffer(nil)
	shClient := exec.Command("sh")
	shClient.Stdin = in
	stdout, err := shClient.StdoutPipe()
	if err != nil {
		return "", err
	}

	if err := shClient.Start(); err != nil {
		return "", err
	}
	for _, cmd := range cmds {
		in.WriteString(cmd + "\n")
		time.Sleep(time.Duration(500) * time.Millisecond)
	}
	in.WriteString("exit\n")

	b, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	if err := shClient.Wait(); err != nil {
		return "", err
	}
	return string(b), err
}
