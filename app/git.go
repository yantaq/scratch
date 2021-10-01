package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

// Clone git repo
func Clone(url, path, targetDir string) {
	// func Clone(url, path, targetPath string) {
	// path := ":identity/aws-access" + ".git"
	// url := "git@ghe.megaleo.com"

	targetDir, err := CreateTemporaryDirectory(targetDir + "-")
	// "/tmp/aws-access"
	if err != nil {
		log.Fatalln("Error creating temp dir: ", err)
	}
	log.Println("Created temp dir: ", targetDir)
	repo := url + path + ".git"
	// if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
	// 	// os.RemoveAll(targetPath)
	// 	// log.Fatalln("target path exist: ", targetPath)
	// }

	// _ = os.Mkdir(targetDir, os.ModeDir)

	// out, err := exec.Command("git", "clone", repo, targetPath).Output()
	cmdStr := "git clone " + repo + " " + targetDir
	out, err := ExecuteCommand(cmdStr)

	if err != nil {
		fmt.Printf("error cloning repo: %s \n %s", repo, out)
		log.Fatal(err)
	}
}

// Branch check branch
func Branch() {
	os.Chdir("/tmp/aws-access")
	out, _ := ExecuteCommand("git branch")
	branches := strings.Fields(out)
	if len(branches) > 2 {
		fmt.Println("more than 1 banch: ", out)
	} else {
		fmt.Println("1 banch: ", branches)
	}
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// ExecuteCommand runs external commands
func ExecuteCommand(commandString string) (string, error) {
	// Tidyup whitespace
	commandString = strings.Join(strings.Fields(commandString), " ")
	parts := strings.SplitAfterN(commandString, " ", 2)
	command := strings.TrimSpace(parts[0])
	argString := parts[1]
	argsList := strings.Split(argString, " ")
	out, err := exec.Command(command, argsList...).Output()
	if err != nil {
		fmt.Printf("error cloning repo: %s \n", out)
		log.Fatal(err)
	}

	return string(out), err
}

// CreateTemporaryDirectory creat temp dir
func CreateTemporaryDirectory(prefix string) (string, error) {
	// Create a local temporary directory
	dirName, err := ioutil.TempDir("/tmp", prefix)
	if err != nil {
		log.Fatalln("Failed to create temporary directory (pattern=)", prefix)
		return "", err
	}
	log.Println("Created temporary directory: ", dirName)
	return dirName, nil
}
