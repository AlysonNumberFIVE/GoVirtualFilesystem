
package main

import (
	"fmt"
	"math/rand"
)

// The main user object.
type user struct {
	userID     uint64         // A randomized integer representing the users's unique ID.
	username   string         // The user's onscreen name.
	accessList map[string]int // A map containing the unique hashes and access rights for each file.
}

// generateRandomID generates a random userID value.
func generateRandomID() uint64 {
	return uint64(rand.Uint32()) << 32 + uint64(rand.Uint32())
}

// createUser creates a user object.
func createUser(username string) *user {
	return &user{
		userID: generateRandomID(),
		username: username,
	}
}

// updateUsername updates the name of the current user.
func (currentUser * user) updateUsername(username string) {
	currentUser.username = username
}
