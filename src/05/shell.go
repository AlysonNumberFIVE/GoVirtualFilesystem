
package main

import (
	"fmt"
	"os"
	"runtime"
	"os/exec"
	"io/ioutil"
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
		if len(segment) == 0 {
			continue 
		}
		if segment == ".." {
			if checker.prev == nil {
				continue 
			}
			checker = checker.prev
		} else if s.doesDirExist(segment, checker) == true {
			checker = checker.directories[segment]
		} else {
			fmt.Printf("Error : %s doesn't exist\n", dirName)
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
	}
	return fs
}

// chDir switches you to a different active directory.
func (s * shell) chDir(dirName string, fs *fileSystem) *fileSystem {
	if dirName == "/" {
		return root
	}
	return s.verifyPath(dirName, fs)
}

func (s * shell) open(filename string, fs *fileSystem) error {

	segments := strings.Split(filename, "/")
	if len(segments) == 1 {
		if _, exists := fs.files[filename]; exists {
			editingFile = fs.files[filename]
			editor()
		} else {
			fmt.Println(filename, ": file doesn't exist.")
		}
	} else {
		dirPath := s.reassemble(segments)
		tmp := s.verifyPath(dirPath, fs)		
		if _, exists := tmp.files[segments[len(segments)-1]]; exists {
			editingFile = tmp.files[segments[len(segments)-1]]
			editor()
			//s.readFile(tmp.files[segments[len(segments)-1]].rootPath)

		} else {
			fmt.Println("cat : file doesn't exist")
		}		
	}
	return nil
}

// reassemble rebuilds the path.
func (s * shell) reassemble(dirPath []string) string {
	counter := 1
	var finishedPath string

	finishedPath = dirPath[0]
	for counter < len(dirPath) - 1 {
		finishedPath = finishedPath + "/" + dirPath[counter]
		counter++
	}
	fmt.Println("reassemble ", finishedPath)
	return finishedPath
}

// readFile reads a file from the filename/path passed in.
func (s * shell) readFile(filename string) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	fmt.Println(string(dat))
} 

// cat dups the contents of a file.
func (s * shell) cat(filename string, fs *fileSystem) {
	var content string

	segments := strings.Split(filename, "/")
	if len(segments) == 1 {
		if _, exists := fs.files[filename]; exists {
			content := string(fs.files[filename].content)
			fmt.Println(content)
		//	s.readFile(fs.files[filename].rootPath)
		} else {
			fmt.Println("cat : file doesn't exist")
		}
	} else {
		dirPath := s.reassemble(segments)
		tmp := s.verifyPath(dirPath, fs)

		if _, exists := tmp.files[segments[len(segments)-1]]; exists {
			content = string(tmp.files[segments[len(segments)-1]].content)
			fmt.Println(content)
			//s.readFile(tmp.files[segments[len(segments)-1]].rootPath)
			fmt.Println("File exists")
		} else {
			fmt.Println("cat : file doesn't exist")
		}
	}
}

// usage prints usage information for shell utilities if used incorrectly.
func (s * shell) usage(comms[] string) bool {
	switch comms[0] {
	case "cd":
		if len(comms) != 2 {
			fmt.Println("Usage : cd [target directory")
			return false
		}
	case "cat":
		if len(comms) != 2 {
			fmt.Println("Usage : cat [target file]")
			return false
		}

	case "open":
		if len(comms) != 2 {
			fmt.Println("Usage : open [file name]")
			return false
		}
	}
	return true
}

// execute contains and runs the commands that are part of the shell.
func (s * shell) execute(comms []string, fs *fileSystem) (*fileSystem, bool) {

	if s.usage(comms) == false {
		return fs, false
	}
	switch comms[0] {
	case "open":
		s.open(comms[1], fs)
	case "cd":
		fs = s.chDir(comms[1], fs)
	case "cat":
		s.cat(comms[1], fs)
	case "clear":
		s.clearScreen()
	default:
		return fs, false
	}
	return fs, true
}
