
package main

import (
	"fmt"
)

/*
** the data structure for each file.
**    name - The name of the file.
**    fileType - The type of the file.
**    content - The file's content in bytes.
**    size - the size in bytes of the file.
*/
type file struct {
	name string
	fileType string
	content bytes
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
	directories []list
	prev *fileSystem
}

/*
** Scans the current directory and builds the VFS from it.
*/
func initFilesystem() {
	fmt.Println("Welcome to the tiny virtual filesystem.")
}

/*
** Resets the VFS and scraps all changes made up to this point.
**  (basically like a rerun of initFilesystem())
*/
func reloadFilesys() {
	fmt.Println("Refreshing...")
}






