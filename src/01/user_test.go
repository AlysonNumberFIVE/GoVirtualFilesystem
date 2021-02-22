
package main

import "testing"

// valueInSlice determines if toFind exists in the IDlist.
func valueInSlice(IDlist []string, toFind string) bool {
	for _, id := range(IDlist) {
		if toFind == id {
			return false
		}
	}
	return true
}

// TestGenerateRandomID tests the functionality of the function
// generateRandomID.
func TestGenerateRandomID(t *testing.T) {
	var IDlist []string
	var newHash string
	counter := 0

	IDlist = append(IDlist, generateRandomID())
	for counter < 10 {
		newHash = generateRandomID()
		if valueInSlice(IDlist, newHash) == false {
			t.Errorf("Hash's aren't being created uniquely.")
		}
		IDlist = append(IDlist, newHash)
		counter++
	}
}

// TestCreateUser tests the functionality of createUser.
func TestCreateUser(t *testing.T) {
	testUser := createUser("AlysonBee")

	if testUser.username != "AlysonBee" {
		t.Errorf("Username wasn't saved.")
	}
	if len(testUser.userID) != 88 {
		t.Errorf("User ID wasn't created accurately.")
	}
	secondUser := createUser("AlysonN")
	if secondUser.userID == testUser.userID {
		t.Errorf("User IDs across 2 user objects are the same.")
	}
}

// TestUpdateUsername tests the functionality of the updateUsername
// method.
func TestUpdateUsername(t *testing.T) {
	currentUsername := "AlysonN"
	testUser := createUser(currentUsername)

	if currentUsername != testUser.username {
		t.Errorf("currentUsername and username don't match.")
	}
	testUser.updateUsername("AlysonV")
	if testUser.username == currentUsername {
		t.Errorf("Username change failed.")
	}
}
