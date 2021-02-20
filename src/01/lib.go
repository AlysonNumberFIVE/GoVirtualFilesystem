
package main

import (
	"fmt"
)

type library struct {
}

/*
** Initialize the library functions.
*/
func initLibrary() *library {
	fmt.Println("Importing library.")
	return &Library{}
}

/*
** This will allow for opening files in virtual space.
*/
func (session * library) open() error {
	fmt.Println("open() called")
	return nil
}

/*
** Close virtual file.
*/
func (session * library) close() error {
	fmt.Println("close() called")
	return nil
}

/*
** Make a virtual directory.
*/
func (session * library) mkDir() error {
	fmt.Println("mkDir() called")
	return nil
}

/*
** Remove a file from the virtual filesystem.
*/
func (session * library) removeFile() error {
	fmt.Println("removeFile() called")
	return nil
}

/*
** Remove a directory from a virtual filesystem.
*/
func (session * library) removeDir() error {
	fmt.Println("removeDir() called")
	return nil
}

/*
** List directory contents.
*/
func (session * library) listDir() error {
	fmt.Println("listDir() called")
	return nil
}