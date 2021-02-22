
package main

import "testing"

// TestInitFilesystem tests initFilesystem.
func TestInitFilesystem(t *testing.T) {
	filesystemTest := initFilesystem()
	newFilesystem := &fileSystem{}

	filesystemTest = newFilesystem
	if root == filesystemTest {
		t.Errorf("root value overwritten.")
	}
}

// TestTearDown tests the functionality of tearDown.
func TestTearDown(t *testing.T) {
	filesystemTest := initFilesystem()
	filesystemTest.tearDown()
}

// TestSaveState tests the functionality of saving state.
func TestSaveState(t *testing.T) {
	filesystemTest := initFilesystem()
	filesystemTest.saveState()
}
