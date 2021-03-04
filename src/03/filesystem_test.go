
package main

import (
	"testing"
	"strings"
)
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

// TestMkDir tests the functionality of the mkdir functionality.
func TestMkDir(t *testing.T) {
	filesystemTest := initFilesystem()
	filesystemTest.mkDir("testDirectory")

	if _, directoryChecker := filesystemTest.directories["testDirectory"]; directoryChecker {
		if filesystemTest.rootPath != "." {
			t.Errorf("Root path not created.")
		}
		newDir := filesystemTest.directories["testDirectory"]
		if newDir.rootPath != "./testDirectory" {
			t.Errorf("testDirectory absolute path not set or correct")
		}
		if newDir.prev != filesystemTest {
			t.Errorf("testDirectory link to parent directory '.' not establised")
		}

	} else {
		t.Errorf("Directory was not successfully created.")		
	}
	if filesystemTest.mkDir("testDirectory") == true {
		t.Errorf("Duplicate file created.")
	}
}

func setupComms(command string) [] string{
	return strings.Split(command, " ")
}

// TestUsage tests the usage() function, ensuring the correct
// usage messages are printed on error.
func TestUsage(t *testing.T) {
	filesystemTest := initFilesystem()
	var comms []string

	// testing mkdir
	comms = setupComms("mkdir TestDirectory")
	if filesystemTest.usage(comms) == false {
		t.Errorf("Usage for mkdir test fails with 1 arg")
	}
	comms = setupComms("mkdir test1 test2 test3 test4")
	if filesystemTest.usage(comms) == false {
		t.Errorf("usage for mkdir test fails with more than 1 argument")
	}
	comms = setupComms("mkdir")
	if filesystemTest.usage(comms) == true {
		t.Errorf("mkdir will create a file you didn't tell it to")
	}

	// testing open
	comms = setupComms("open file1 file2 file3")
	if filesystemTest.usage(comms) == true {
		t.Errorf("Opening more than 1 file shouldn't be possible")
	}
	comms = setupComms("open")
	if filesystemTest.usage(comms) == true {
		t.Errorf("You can't open an unspecified file")
	}
	comms = setupComms("open file1")
	if filesystemTest.usage(comms) == false {
		t.Errorf("Opening a target file failed")
	}
}
