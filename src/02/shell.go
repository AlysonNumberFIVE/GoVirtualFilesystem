
package main

import (
	"strings"
)

// shellLoop runs the main shell loop for the filesystem.
func shellLoop(currentUser *user) {

	filesystem := initFilesystem()
	prompt := currentUser.initPrompt()
	for {
		input, _ := prompt.Readline()
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			continue 
		}
		filesystem.execute(input)
	}
}

func main() {
	currentUser := initUser()

	shellLoop(currentUser)
}
