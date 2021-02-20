
package main

import (
	"fmt"
)

// A global list of all files created and their respective names for
// ease of lookup.
var globalFileTable map[uint64]string

/*
** The data structure for each file.
**    name - The name of the file.
**    pathFromRoot - The absolute path of the file
**    fileHash - The unique file hash assigned to this file on
**        creation
**    fileType - The type of the file.
**    content - The file's content in bytes.
**    size - the size in bytes of the file.
*/
type file struct {
	name string
	pathFromRoot string
	fileHash uint64
	fileType string
	content byte
	size uint64
}

/*
** The core struct that makes up the filesystem's file/directory
**      structure.
**    directory - The name of the current directory we're in.
**    files - The list of files in this directory.
**    directories - The list of directories in this directory.
**    prev - a reference pointer to this directory's parent directory.
*/

type fileSystem struct {
	directory string
	files []file
	directories []fileSystem
	prev *fileSystem
}

// Root node.
var root *fileSystem	

/*
** Scans the current directory and builds the VFS from it.
*/
func initFilesystem() * fileSystem {
	// recursively grab all files and directories from this level downwards.
	fmt.Println("Welcome to the tiny virtual filesystem.")
	return root
}

/*
** Resets the VFS and scraps all changes made up to this point.
**  (basically like a rerun of initFilesystem())
*/
func (root * fileSystem) reloadFilesys() {
	fmt.Println("Refreshing...")
}

/*
** Gracefully ends the current session.
*/
func (root * fileSystem) tearDown() {
	fmt.Println("Teardown")
}

/*
** Saves the state of the VFS at this time.
*/
func (root * fileSystem) saveState() {
	fmt.Println("Save the current state of the VFS")
}
