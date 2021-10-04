package app

import (
	"log"

	"github.com/yantaq/scratch/cmd"
)

// Open opens a broswer page with give url
func Open(url string) {
	var out string
	var err error
	if len(url) < 1 {
		cmdStr := "open https://google.com"
		out, err = cmd.Run(cmdStr)
	} else {
		cmdStr := "open " + url
		out, err = cmd.Run(cmdStr)
	}

	if err != nil {
		log.Fatalln("Error: ", out, err.Error())
	}
}
