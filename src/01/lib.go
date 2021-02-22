
package main

import (
	"fmt"
)

type library struct {
}

// initLibrary initializes the library functions.
func initLibrary() *library {
	fmt.Println("Importing library.")
	return &library{}
}

// open will allow for opening files in virtual space.
func (session * library) open() error {
	fmt.Println("open() called")
	return nil
}

// close closes open virtual files.
func (session * library) close() error {
	fmt.Println("close() called")
	return nil
}

// mkDir makes a virtual directory.
func (session * library) mkDir() error {
	fmt.Println("mkDir() called")
	return nil
}

// removeFile removes a file from the virtual filesystem.
func (session * library) removeFile() error {
	fmt.Println("removeFile() called")
	return nil
}

// removeDir removes a directory from the virtual filesystem.
func (session * library) removeDir() error {
	fmt.Println("removeDir() called")
	return nil
}

// listDir lists a directory's contents.
func (session * library) listDir() error {
	fmt.Println("listDir() called")
	return nil
}