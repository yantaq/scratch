package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Clone git repo into target dir
func Clone(url, path, targetDir string) (string, error) {
	targetDir, err := CreateTemporaryDirectory(targetDir + "-")
	if err != nil {
		log.Fatalln("Error creating temp dir: ", err)
		return "", err
	}
	repo := url + path
	// if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
	// os.RemoveAll(targetPath)
	// log.Fatalln("target path exist: ", targetPath)
	// }
	// _ = os.Mkdir(targetPath, os.ModeDir)

	// out, err := exec.Command("git", "clone", repo, targetPath).Output()
	cmdStr := "git clone " + repo + " " + targetDir
	out, err := ExecuteCommand(cmdStr)

	if err != nil {
		log.Fatal("Error cloning repo: ", repo, out, err)
		return "", err
	}

	return targetDir, nil
}

// ExecuteCommand runs external commands
func ExecuteCommand(commandString string) (string, error) {
	// Tidyup whitespace
	commandString = strings.Join(strings.Fields(commandString), " ")
	parts := strings.SplitAfterN(commandString, " ", 2)
	command := strings.TrimSpace(parts[0])
	argString := parts[1]
	argsList := strings.Split(argString, " ")
	log.Println("running cmd: ", command, argString)
	out, err := exec.Command(command, argsList...).Output()
	if err != nil {
		fmt.Println("Error: ", command, argString)
		log.Fatal(err)
	}

	return string(out), err
}

// CreateTemporaryDirectory create temp dir
func CreateTemporaryDirectory(prefix string) (string, error) {
	// Create a local temporary directory
	dirName, err := ioutil.TempDir("/tmp", prefix)
	if err != nil {
		log.Fatalln("Failed to create temporary directory pattern=", prefix)
		return "", err
	}
	log.Println("Created temporary directory: ", dirName)
	return dirName, nil
}

// CreatePushBranch branch check branch
func CreatePushBranch() {
	employees := []string{"jon.doe"}
	projDir := "/tmp/proj dir"
	os.Chdir(projDir)
	err := os.Chdir(projDir)
	if err != nil {
		log.Fatalln("Error changing dir: ", err)
	}

	name := strings.Replace(employees[0], ".", "_", 1)
	branchName := "remove_" + name
	createBranchCmd := "git checkout -b " + branchName
	out, _ := ExecuteCommand(createBranchCmd)
	fmt.Println("Branch created: ", branchName, out)

	checkoutBranchCmd := "git checkout " + branchName
	out, _ = ExecuteCommand(checkoutBranchCmd)
	fmt.Println("Branch checkout: ", branchName, out)

	commitMsg := "remove_" + name
	commitBranchCmd := "git commit -a -m " + commitMsg
	log.Println(commitBranchCmd)
	out, _ = ExecuteCommand(commitBranchCmd)

	pushUpstreamCmd := "git push --set-upstream origin " + branchName
	out, _ = ExecuteCommand(pushUpstreamCmd)
	log.Println(pushUpstreamCmd)
	if strings.Contains(out, "pull/new") {
		log.Println("pushed to upstream: ", out)
	}
}
