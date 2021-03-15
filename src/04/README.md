# Ex 04 - Completing Filesystem Initialization

The accompanying article for this step can be found <a href="">here</a>.<br><br>
This week, work is being done on filesystem initialization and the handling of individual files. The goal is to make a digital copy of the current directory we're in and to implement the `cat` command.<br>
Here's the eventual goal for this week:<br><br>
<img src="https://github.com/AlysonBee/GoVirtualFilesystem/blob/master/assets/demo3.gif"  height="400" />

## Completing Initialization

The code can be found in `filesystem.go`<br><br>

Initialization recursively creates a copy of all the files and directories from the current directory you start the filesystem from and everything downwards.<vr>
Overall functionality makes use of two functions; `makeFilesystem` and `testFilesystemCreation` (this name was a placeholder but will have to be dropped... naturally).
```
func makeFilesystem(dirName string, rootPath string, prev *fileSystem) * fileSystem {
	return &fileSystem{
		name: dirName,
		rootPath: rootPath,
		files: make(map[string]*file),
		directories: make(map[string]*fileSystem),
		prev: prev,
	}
}
```

```
func testFilesystemCreation(dirName string, fs *fileSystem) *fileSystem{
	var fi os.FileInfo
	var fileName os.FileInfo

	if dirName == "." {
		root = makeFilesystem(".", ".", nil)
		fs = root
	}
	index := 0
	files, _ := ioutil.ReadDir(dirName)
	for index < len(files) {
		fileName = files[index]
		fi, _ = os.Stat(dirName + "\\" + fileName.Name())
		mode := fi.Mode()
		if mode.IsDir() {
			fs.directories[fileName.Name()] = makeFilesystem(fileName.Name(), dirName + "/" + fileName.Name(), fs)
			testFilesystemCreation(dirName + "\\" + fileName.Name(), fs.directories[fileName.Name()])
		} else {
			fs.files[fileName.Name()] = &file{
				name: fileName.Name(),
				rootPath: dirName + "/" + fileName.Name(),
			}
		}
		index++
	}
	return fs
}
```
`makeFilesystem` is an abstraction for creating a single `fileSystem` object. The overall logic of the function is that if a file is detected, it's added to the `files ` list in the current`fileSystem` object we're in. And if it's a new directory, a recursive call is made on `testCreateFilesystem` and the process recursivel continues.<br>
`os.Stat` allows us to pull information on files/directories. For now, that info is only used to determine if a file is a directory or not.<BR><BR>
A key implementation detail of this function is that entire files aren't stored in the virtual filesystem copy. Only the paths to the files. The idea being that only files that are eventually modified are going to be stored to save on space and iniialization time, although this might be adjusted based on how much sense (or lack thereof) this makes later on.


## Reading Files - The CAT command.

Open the file `shell.go`.<br><br>

The `cat` command prints out the contents of the file that we target.<br><br>
```
func (s * shell) cat(filename string, fs *fileSystem) {
	segments := strings.Split(filename, "/")
	if len(segments) == 1 {
		if _, exists := fs.files[filename]; exists {
			s.readFile(fs.files[filename].rootPath)
		} else {
			fmt.Println("cat : file doesn't exist")
		}
	} else {
		dirPath := s.reassemble(segments)
		tmp := s.verifyPath(dirPath, fs)

		if _, exists := tmp.files[segments[len(segments)-1]]; exists {
			s.readFile(tmp.files[segments[len(segments)-1]].rootPath)
			fmt.Println("File exists")
		} else {
			fmt.Println("cat : file doesn't exist")
		}
	}
}
```

The logic of the `cat` command makes use of the `verifyPath` method to ensure the path being passed in is valid if a path exists. The `files` array is then checked to determine if the target file exists. As we've only handled reading and not writing at this point, the `cat` command dumps info from the native filesystem using the `rootPath` as reference.<br>
Once writing is implemented, a check will be added to determine if `cat` should read from either the native filesystem or the `contents` field of the `file` object.













