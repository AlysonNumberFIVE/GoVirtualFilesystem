
package main

import (
	"strings"
)

// shellLoop runs the main shell loop for the filesystem.
func shellLoop(currentUser *user) {
	var shellFlag bool

	shell := initShell()
	fs := initFilesystem()
	prompt := currentUser.initPrompt()
	for {
		input, _ := prompt.Readline()
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			continue 
		}

		comms := strings.Split(input, " ")

		fs, shellFlag = shell.execute(comms, fs)
		if shellFlag == true {
			continue 
		}

		fs, shellFlag = fs.execute(comms)
		if shellFlag == true {
			continue 
		}
	}
}

func main() {
	currentUser := initUser()

	shellLoop(currentUser)
}



