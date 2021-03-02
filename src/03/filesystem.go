
package main

import (
	"fmt"
	"os"
)

// A global list of all files created and their respective names for
// ease of lookup.
var globalFileTable map[uint64]string

// The data structure for each file.
type file struct {
	name         string // The name of the file.
	rootPath     string // The absolute path of the file.
	fileHash     uint64 // The unique file hash assigned to this file on creation.
	fileType     string // The type of the file.
	content      []byte // The file's content in bytes.
	size         uint64 // The size in bytes of the file.
}

//  The core struct that makes up the filesystem's file/directory
type fileSystem struct {
	name	    string                 // The name of the current directory we're in.
	rootPath	string                 // The absolute path to this directory.
	files       []file                 // The list of files in this directory.
	directories map[string]*fileSystem // The list of directories in this directory.
	prev        *fileSystem            // a reference pointer to this directory's parent directory.	
}

// Root node.
var root *fileSystem	

// initFilesystem scans the current directory and builds the VFS from it.
func initFilesystem() * fileSystem {
	// recursively grab all files and directories from this level downwards.
	root = &fileSystem{
		name: ".",
		rootPath: ".",
		directories: make(map[string]*fileSystem),
	}
	fs := root
	fmt.Println("Welcome to the tiny virtual filesystem.")
	return fs
}

// pwd prints the current working directory.
func (fs * fileSystem) pwd() {
	fmt.Println(fs.rootPath)
}

// reloadFilesys Resets the VFS and scraps all changes made up to this point.
// (basically like a rerun of initFilesystem())
func (fs  * fileSystem) reloadFilesys() {
	fmt.Println("Refreshing...")
}

// tearDown gracefully ends the current session.
func (fs  * fileSystem) tearDown() {
	fmt.Println("Teardown")
}

// saveState aves the state of the VFS at this time.
func (fs  * fileSystem) saveState() {
	fmt.Println("Save the current state of the VFS")
}

// open will allow for opening files in virtual space.
func (fs  * fileSystem) open() error {
	fmt.Println("open() called")
	return nil
}

// close closes open virtual files.
func (fs  * fileSystem) close() error {
	fmt.Println("close() called")
	return nil
}

// mkDir makes a virtual directory.
func (fs  * fileSystem) mkDir(dirName string) error {
	newDir := &fileSystem{
		name: dirName,
		rootPath: fs.rootPath + "/" + dirName,
		directories: make(map[string]*fileSystem),
		prev: fs,
	}
	fs.directories[dirName] = newDir
	return nil
}

// removeFile removes a file from the virtual filesystem.
func (fs  * fileSystem) removeFile() error {
	fmt.Println("removeFile() called")
	return nil
}

// removeDir removes a directory from the virtual filesystem.
func (fs  * fileSystem) removeDir() error {
	fmt.Println("removeDir() called")
	return nil
}

// listDir lists a directory's contents.
func (fs  * fileSystem) listDir() {

	if fs.files != nil {
		fmt.Println("File:")
		for _, file := range fs.files {
			fmt.Printf("\t%s\n", file.name)
		}
	}
	if len(fs.directories) > 0 {
		fmt.Println("Directories:")
		for dirName := range fs.directories {
			fmt.Printf("\t%s\n", dirName)
		}
	}
}

func (fs * fileSystem) usage(comms []string) bool {
	switch comms[0] {
	case "mkdir":
		if len(comms) < 2 {
			fmt.Println("Usage : mkdir [list of directories to make]")
			return false
		}
	case "open":
		if len(comms) != 2 {
			fmt.Println("Usage : open [file name]")
			return false
		}
	case "rm":
		if len(comms) != 2 {
			fmt.Println("Usage : rm [list of files]")
			return false
		}
	}
	return true
}

// execute runs the commands passed into it.
func (fs * fileSystem) execute(comms []string) * fileSystem {
	if fs.usage(comms) == false {
		return fs
	}
	switch comms[0] {
	case "mkdir":
		fs.mkDir(comms[1])
	case "pwd":
		fs.pwd()
	case "open":
		fs.open()
	case "close":	
		fs.close()
	case "ls":
		fs.listDir()
	case "rm":
		fs.removeFile()
		fs.removeDir() 
	case "exit":
		fs.tearDown()
		os.Exit(1)
	default:
		fmt.Println(comms[0], ": Command not found")
	}
	return fs
}
