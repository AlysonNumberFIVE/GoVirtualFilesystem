
package main

import (
	"strings"
	"fmt"
)

func help() {
	fmt.Println("ls              : list files and directories")
	fmt.Println("cd              : change directory")
	fmt.Println("clear           : clear screen")
	fmt.Println("open [filename] : open the specified filename.")
	fmt.Println("pwd             : print working directory")
	fmt.Println("cat             : dump the contents of a file.")
	fmt.Println("exit            : exit the shell/filesystem.")
}

// shellLoop runs the main shell loop for the filesystem.
func shellLoop(currentUser *user) {
	var shellFlag bool

	shell := initShell()
	fs := initFilesystem()
	prompt := currentUser.initPrompt()
	help()
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



