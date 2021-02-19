
package main

import (
	"fmt"
	"bufio"
	"os"
)

/*
** The main shell loop for the filesystem.
*/
func ShellLoop() {

	library := InitLibrary()
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
				library.Open()
			case "close":
				library.Close()
			case "remove":
				library.RemoveDir()
			case "ls":
				library.listDir()
			default:
				fmt.Println(input, ": Command not found")
		}
	}
}


func main() {

	ShellLoop()
}
