
package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
	"strings"
)

type shell struct {
	env map[string]string
}

func initShell() *shell {
	thisPath, _ := os.Getwd()
	rootPath := strings.Replace(thisPath, "\\", "/", -1)
	env["PWD"] = rootPath 	
	env["HOME"]  = rootPath	
	env["OLDPWD"] = rootPath
	return &shell{
		env: env,
	}
}

// Env variable stores current directory infomration.
var env map[string]string


// clearScreen clears the terminal screen.
func (s *shell) clearScreen() {
	clear := make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func () {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	cls, ok := clear[runtime.GOOS]
	if ok {
		cls()
	}
}

// doesDirExist checks if the dirName directory exists.
func (s *shell) doesDirExist(dirName string, fs *fileSystem) bool {
	fmt.Println("DoesExist entered")
	if _, found := fs.directories[dirName]; found {
		return true
	}
	return false
}

// chDir lists a directory's contents.
func (s * shell) chDir(dirName string, fs *fileSystem) *fileSystem {
	checker := fs
	segments := strings.Split(dirName, "/")

	for _, segment := range segments {
		if segment == ".." {
			if checker.prev == nil {
				break
			}
			checker = checker.prev
		} else if s.doesDirExist(segment, checker) == true {
			checker = checker.directories[segment]
		} else {
			fmt.Printf("Error : %s doesn't exist", dirName)
			return fs
		}
	}
	return checker
}