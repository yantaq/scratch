package app

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// Clone git repo
func Clone() {
	// func Clone(url, path, targetPath string) {
	path := ":identity/aws-access" + ".git"
	url := "git@ghe.megaleo.com"
	targetPath := "/tmp/aws-access"
	repo := url + path
	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		// os.RemoveAll(targetPath)
		// log.Fatalln("target path exist: ", targetPath)
	}

	_ = os.Mkdir(targetPath, os.ModeDir)

	out, err := exec.Command("git", "clone", repo, targetPath).Output()
	if err != nil {
		fmt.Printf("error cloning repo: %s \n %s", repo, out)
		log.Fatal(err)
	}
}

func Branch() {

}
