// +build windows

package main

import (
	"path/filepath"
	"regexp"
	"syscall"
)

const osHaveSigTerm = false

func envNixToWin(line []byte) string {
	nr := regexp.MustCompile("\\${?(\\w+)}?")
	out := nr.ReplaceAll(line, []byte("%$1%"))
	return string(out)
}

func ShellInvocationCommand(interactive bool, root, command string) []string {
	command = filepath.FromSlash(command)  // In case we have unix style seperators
	command = envNixToWin([]byte(command)) // In case we get unix style env vars
	return []string{"cmd", "/C", command}
}

func (p *Process) PlatformSpecificInit() {
	// NOP on windows for now.
	return
}

func (p *Process) SendSigTerm() {
	panic("SendSigTerm() not implemented on this platform")
}

func (p *Process) SendSigKill() {
	p.Signal(syscall.SIGKILL)
}
