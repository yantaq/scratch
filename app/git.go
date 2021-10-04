package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/yantaq/scratch/cmd"
)

// Clone git repo into target dir
func Clone(url, path, targetDir string) (string, error) {
	targetDir, err := CreateTempDir(targetDir + "-")
	if err != nil {
		log.Fatalln("Error creating temp dir: ", err)
		return "", err
	}
	repo := url + "/" + path
	cmdStr := "git clone " + repo + " " + targetDir
	out, err := cmd.Run(cmdStr)
	if err != nil {
		log.Fatalln("Error: ", out, err.Error())
	}

	path, err = os.Getwd()
	if err != nil {
		log.Println(err)
	}
	log.Println(path)

	lsCmd := "ls -l ."
	out, err = cmd.Run(lsCmd)

	if err == nil {
		log.Println(out)
	}

	return targetDir, nil
}

// CreateTempDir create temp dir
func CreateTempDir(prefix string) (string, error) {
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
	out, _ := cmd.Run(createBranchCmd)
	fmt.Println("Branch created: ", branchName, out)

	checkoutBranchCmd := "git checkout " + branchName
	out, _ = cmd.Run(checkoutBranchCmd)
	fmt.Println("Branch checkout: ", branchName, out)

	commitMsg := "remove_" + name
	commitBranchCmd := "git commit -a -m " + commitMsg
	log.Println(commitBranchCmd)
	out, _ = cmd.Run(commitBranchCmd)

	pushUpstreamCmd := "git push --set-upstream origin " + branchName
	out, _ = cmd.Run(pushUpstreamCmd)
	log.Println(pushUpstreamCmd)
	if strings.Contains(out, "pull/new") {
		log.Println("pushed to upstream: ", out)
	}
}
