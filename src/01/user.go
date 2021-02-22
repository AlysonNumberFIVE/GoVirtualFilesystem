
package main

import (
	"crypto/rand"
	"encoding/base64"
)

// The main user object.
type user struct {
	userID     string         // A randomized hash string representing the users's unique ID.
	username   string         // The user's onscreen name.
	accessList map[string]int // A map containing the unique hashes and access rights for each file.
}

// generateRandomID generates a random userID value.
func generateRandomID() string {
	// return uint64(rand.Uint32()) << 32 + uint64(rand.Uint32())
    bytes := make([]byte, 64)
   	rand.Read(bytes)
    return base64.URLEncoding.EncodeToString(bytes)
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
