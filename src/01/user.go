
package main

import (
	"fmt"
	"math/rand"
)

/*
** User Object
** 		userID - A randomized integer representing the users's unique ID
** 		username - The user's onscreen name
**		accessList - A map containing the unique hashes of existing files
**				and any custom/specific access rights added
**				by the user.
*/
type user struct {
	userID uint64
	username string
	accessList map[string]int
}

/*
** Generates a random userID.
*/
func generateRandomID() uint64 {
	return uint64(rand.Uint32()) << 32 + uint64(rand.Uint32())
}


/*
** Creates a user object.
*/
func createUser(username string) *user {
	return &user{
		userID: generateRandomID(),
		username: username,
	}
}

/*
** Update the name of the current user.
*/
func (currentUser * user) updateUsername(username string) {
	currentUser.username = username
}



func main() {
	newUser := createUser("AlysonV")
	fmt.Println("new user logged in : ", newUser.username)
	fmt.Println("User ID : ", newUser.userID)
	newUser.updateUsername("VivianS")
	fmt.Println("new user is ", newUser.username)
	fmt.Println("User ID : ", newUser.userID)
}







