
package main

import (
	"fmt"
	"testing"
)

// TestVerifyPath determines that the verifyPath function catches
// valid and invalid paths accurately.
func TestVerifyPath(t *testing.T) {
	filesystemTest := initFilesystem()
	shell := initShell()

	filesystemTest.mkDir("dir1")
	filesystemTest.mkDir("dir2")
	filesystemTest = shell.chDir("dir1", filesystemTest)
	filesystemTest.mkDir("dir3")
	filesystemTest = shell.chDir("dir3", filesystemTest)
	filesystemTest.mkDir("dir4")
	filesystemTest = shell.chDir("dir4", filesystemTest)

	// currently in ./dir1/dir3/dir4
	filesystemTest = shell.chDir("../..", filesystemTest)
	if filesystemTest.rootPath != "./dir1" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Inaccuracy in changing directories")
	}
	filesystemTest = shell.chDir("dir3/dir4", filesystemTest)
	if filesystemTest.rootPath != "./dir1/dir3/dir4" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Inaccuracy in changing directories")
	}

	filesystemTest.mkDir("helloWorld")
	filesystemTest.mkDir("goodByeUniverse")
	// currently in ./dir1/dir3/dir4/helloWord
	filesystemTest = shell.chDir("helloWorld", filesystemTest)
	if filesystemTest.rootPath != "./dir1/dir3/dir4/helloWorld" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Inaccuracy in changing directories")
	}
	// currently in ./dir2
	filesystemTest = shell.chDir("../../dir4/../../../dir2", filesystemTest)
	if filesystemTest.rootPath != "./dir2" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Inaccuracy in changing directories")
	} 

	// currently in ./dir1/dir3/dir4/goodByeUniverse
	filesystemTest = shell.chDir("../dir1/dir3/dir4/goodByeUniverse", filesystemTest)
	if filesystemTest.rootPath != "./dir1/dir3/dir4/goodByeUniverse" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Inaccuracy in changing directories")
	}

	filesystemTest = shell.chDir("../../dir4/../../dir2", filesystemTest)
	if filesystemTest.rootPath != "./dir1/dir3/dir4/goodByeUniverse" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Directory wasn't supposed to change")
	}

	filesystemTest = shell.chDir("/", filesystemTest)
	if filesystemTest.rootPath != "." {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Navigation to root doesn't work")
	}

	filesystemTest = shell.chDir("dir2", filesystemTest)
	filesystemTest = shell.chDir("/dir1/dir3/dir4/goodByeUniverse", filesystemTest)
	// we're at /dir1/dir3dir4/helloWorld
	if filesystemTest.rootPath != "./dir1/dir3/dir4/goodByeUniverse" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Navigation from root doesn't exist")
	}

	filesystemTest = shell.chDir("../../", filesystemTest)
	if filesystemTest.rootPath != "./dir1/dir3" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Navigation doesn't work with string ending with /")
	}

	filesystemTest = shell.chDir("aosdihgoiahsdgahsdgiahsga", filesystemTest)
	if filesystemTest.rootPath != "./dir1/dir3" {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Nonexistent files don't navigate")
	}

	filesystemTest = shell.chDir("../../../../../../../../../../", filesystemTest)
	if filesystemTest.rootPath != "." {
		fmt.Println(filesystemTest.rootPath)
		t.Errorf("Infinite backwards traversal doesn't work")
	}
}

// TestDoesDirExist tests if doesDirExist works.
func TestDoesDirExist(t *testing.T) {
	filesystemTest := initFilesystem()
	shell := initShell()

	filesystemTest.mkDir("newDirectory")
	if shell.doesDirExist("newDirectory", filesystemTest) == false {
		t.Errorf("Directory falsly flagged as nonexistent")
	}
	if shell.doesDirExist("doesNotExist", filesystemTest) == true {
		t.Errorf("Directory falsly flagged as existing")
	}
}

// TestHandleRootNav tests that handleRootNav returns us to root
// upon passing in a path that starts with a '/' character.
func TestHandleRootNav(t *testing.T) {
	filesystemTest := initFilesystem()
	shell := initShell()

	filesystemTest.mkDir("dir1")
	filesystemTest = shell.chDir("dir1", filesystemTest)
	filesystemTest.mkDir("dir2")
	filesystemTest = shell.chDir("dir2", filesystemTest)
	filesystemTest.mkDir("dir3")
	filesystemTest = shell.chDir("dir3", filesystemTest)

	tmpFs := shell.handleRootNav("/root/path", filesystemTest)
	if tmpFs.rootPath != "." {
		t.Errorf("Changing to root isn't working")
	}
	tmpFs = shell.handleRootNav("..", filesystemTest)
	if tmpFs != filesystemTest {
		t.Errorf("Handle root not working")
	}
}

// TestChDir tests that chDir behaves as expected.
// Only existent and nonexist directories will be checked
// as path correction is checked by TestVerifyPath.
func TestChDir(t *testing.T) {
	filesystemTest := initFilesystem()
	shell := initShell()

	
}











