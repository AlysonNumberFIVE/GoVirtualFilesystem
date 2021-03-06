
package main

import (
	"fmt"
	"strings"
)

// shellLoop runs the main shell loop for the filesystem.
func shellLoop(currentUser *user) {

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
		if comms[0] == "cd" {
			if len(comms) != 2 {
				fmt.Println("Usage : cd [directory]")
				continue 
			}
			fs = shell.chDir(comms[1], fs)
		} else if comms[0] == "clear" {
			shell.clearScreen()
		} else {
			fs.execute(comms)
		}
	}
}

func main() {

	currentUser := initUser()

	shellLoop(currentUser)
}



