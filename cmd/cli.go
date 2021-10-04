package cmd

import (
	"os/exec"
	"strings"
)

// Run execute external command
func Run(cmdStr string) (string, error) {
	cmdStr = strings.Join(strings.Fields(cmdStr), " ")
	s := strings.SplitAfterN(cmdStr, " ", 2)
	cmd := strings.TrimSpace(s[0])
	arg := s[1]
	args := strings.Split(arg, " ")
	out, err := exec.Command(cmd, args...).CombinedOutput()
	return string(out), err
}
