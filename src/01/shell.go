
package main

import (
	"fmt"
	"bufio"
	"os"
)

// shellLoop runs the main shell loop for the filesystem.
func shellLoop() {

	library := initLibrary()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("$>")
		input, _ := reader.ReadString('\n')
	
		if input == "\r\n" {
			continue 
		}

		input = input[:len(input) - 2]

		switch input {	
		case "open":
			library.open()
		case "close":
			library.close()
		case "remove":
			library.removeDir()
		case "ls":
			library.listDir()
		default:
			fmt.Println(input, ": Command not found")
		}
	}
}