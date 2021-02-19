
package main

import (
	"fmt"
)

type Library struct {
}

/*
** Initialize the library functions.
*/
func InitLibrary() *Library {
	fmt.Println("Importing library.")
	return &Library{}
}

/*
** This will allow for opening files in virtual space.
*/
func (session * Library) Open() error {
	fmt.Println("Open() called")
	return nil
}

/*
** Close virtual file.
*/
func (session * Library) Close() error {
	fmt.Println("Close() called")
	return nil
}

/*
** Make a virtual directory.
*/
func (session * Library) MkDir() error {
	fmt.Println("MkDir() called")
	return nil
}

/*
** Remove a file from the virtual filesystem.
*/
func (session * Library) RemoveFile() error {
	fmt.Println("RemoveFile() called")
	return nil
}

/*
** Remove a directory from a virtual filesystem.
*/
func (session * Library) RemoveDir() error {
	fmt.Println("RemoveDir() called")
	return nil
}

/*
** List directory contents.
*/
func (session * Library) listDir() error {
	fmt.Println("ListDir() called")
	return nil
}

