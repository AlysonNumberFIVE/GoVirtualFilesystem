
package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
	"strings"
)

// our shell object.
type shell struct {
	env map[string]string // the environment varialbes.
}

// Env variable stores current directory infomration.
var env map[string]string

// initShell initializes our shell object.
func initShell() *shell {
	env = make(map[string]string)
	return &shell{
		env: env,
	}
}

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
	if _, found := fs.directories[dirName]; found {
		return true
	}
	return false
}

// verifyPath ensures that the path in dirName exists.
func (s * shell) verifyPath(dirName string, fs *fileSystem) *fileSystem {

	checker := s.handleRootNav(dirName, fs)
	segments := strings.Split(dirName, "/")

	for _, segment := range segments {
		if segment == ".." {
			if checker.prev == nil {
				continue 
			}
			checker = checker.prev
		} else if s.doesDirExist(segment, checker) == true {
			checker = checker.directories[segment]
		} else {
			fmt.Println("Doesn't exist")
			fmt.Printf("Error : %s doesn't exist", dirName)
			return fs
		}
	}
	return checker 
}

// handleRootNav determines if we'll be handling changing directory
// starting from our root.
func (s * shell) handleRootNav(dirName string, fs *fileSystem) *fileSystem {

	if dirName[0] == '/' {
		return root
	} else {
		return fs
	}
}

// chDir switches you to a different active directory.
func (s * shell) chDir(dirName string, fs *fileSystem) *fileSystem {

	if dirName == "/" {
		return root
	}

	return s.verifyPath(dirName, fs)
}