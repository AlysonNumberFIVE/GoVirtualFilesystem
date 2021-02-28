
package main

import (
	"fmt"
	"log"
	"github.com/chzyer/readline"
)

// setName gets a custom username from the current user.
func setName() string {

	var username string
	line, err := readline.New(">")
	if err != nil {
		log.Fatal(err)
	}

	for {
		fmt.Println("Please enter a username (1 for Anonymous):")
		input, err := line.Readline()
		if err != nil {
			log.Fatal(err)
		}
		if input == "1" {
			fmt.Println("Anonymous it is")
			username = "Anon"
			break 
		}
		if len(input) > 2 {
			fmt.Println("Welcome ", input)
			username = input
			break 
		}
	}
	return username
}

// initUser initializes the user object on startup.
func initUser() *user {
	username := setName()
	currentUser := createUser(username)
	return currentUser
}
