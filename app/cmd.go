package app

import (
	"fmt"
	"os"
	"os/exec"
)

// CmdOutput output external command run output
func CmdOutput() {
	os.Chdir("/tmp/project dir")
	output, err := exec.Command("git", "checkout", "-b", "branch_name").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Println(string(output))
}
