# Ex 03 - Implementing a Shell.

The accompanying article for this week can be found <a href="">here</a><br><br>
The plan changed from last week and fleshing out the interactive shell became a necessity. The goal for this week is to have a shell with the following behavior:<br>
<img src="https://github.com/AlysonBee/GoVirtualFilesystem/blob/master/assets/demo2.gif"  height="400" />
## chDir and mkDir
Let's start with <b>chDir</b><br>
Open `shell.go`<BR><br>
The change directory function was moved to a completely new object called `shell`
```
type shell struct {
	env map[string]string // the environment varialbes.
}
```
This change happened when I had trouble with making the directory change work from the filesystem object so I made it a part of the shell. This was also helped by teh fact that most shell prompts have `chDir` (`cd`) built into the (see this <a href="https://en.wikipedia.org/wiki/Cd_(command)">Wiki</a> I read on the topic).<br>
chDir is built around a helper function that validates the existence of a valid path passed into it.
`chDir`:
```
func (s * shell) chDir(dirName string, fs *fileSystem) *fileSystem {
	if dirName == "/" {
		return root
	}
	return s.verifyPath(dirName, fs)
}
```
`verifyPath`:
```
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
```

`verifyPath` loops through each directory one by one to find the directory targeted in the `dirName` variable. If a `..` is detected, a move up to the parent directory is done with the exception of being at root, where nothing will happen.<br>
Starting off a path string with `/` will navigate you to the root directory and then begin navigation from there (i.e `cd /`, `cd /hello` etc.).<br>
A small design change in the `fileSystem` struct was made; changing the `directories` member from an array of structs to a map; the key being the directory name, the value being the directory struct associated with it, to improve access speed.<br>
Before:
```
type fileSystem struct {
        directory   string       // The name of the current directory we're in.
        files       []file       // The list of files in this directory.
        directories []fileSystem // The list of directories in this directory.
        prev        *fileSystem  // a reference pointer to this directory's parent directory.
}
```
After:
```
type fileSystem struct {
        name        string                 // The name of the current directory we're in.
        rootPath    string                 // The absolute path to this directory.
        files       []file                 // The list of files in this directory.
        directories map[string]*fileSystem // The list of directories in this directory.
        prev        *fileSystem            // a reference pointer to this directory's parent directory.
}
```
While talking about changes made to `fileSystem`, you'll see that a `rootPath` variable is added. This just holds the absolute path of the current directory. Its incredibly useful for the function `pwd` which is just a one-liner.
```
func (fs * fileSystem) pwd() {
	fmt.Println(fs.rootPath)
}
```
This nicely brings us to our next function; <b>mkDir</b><br><br>
```
func (fs  * fileSystem) mkDir(dirName string) bool {

	if _, exists := fs.directories[dirName]; exists {
		fmt.Println("mkdir : directory already exists")
		return false
	}

	newDir := &fileSystem{
		name: dirName,
		rootPath: fs.rootPath + "/" + dirName,
		directories: make(map[string]*fileSystem),
		prev: fs,
	}
	fs.directories[dirName] = newDir
	return false
}
```
You'll see that `rootPath` generates a new root path by adding the newly created filename to the current `rootPath` value of the directory we're in. But before creating anything, we check for duplicate names.<BR>


