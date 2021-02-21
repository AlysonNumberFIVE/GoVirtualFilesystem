# Ex01 - Laying the Foundation

This is a shortened version of <a href="https://alysonn.medium.com/a-virtual-filesystem-in-go-creating-our-foundation-9af62b0e82db">this</a> article in exercise form.<br><br>
The first exercise entails creating the groundwork for the upcoming projects. The steps below are as follows<br><br>

### 1. Download Go

The first step is to setup your machine to ensure you're able to use Go. You can download the binaries <a href="https://golang.org/doc/install">at the official site</a> for either Windows, MacOS or Linux. Unzip the binary and follow the prompts to install.

### 2. Setup your project

To setup your directory, you can simply stick to this directory's layout of the four `.go` files. You can then create a `go.mod` to keep track of your dependencies and modules by running
```
go mod init GoVFS
```
You can replace GoVFS with whatever name you'd prefer to name your project.

### 3. An Overview

The filesystem will be divided into 3 main components; the user object, the filesystem itself and the lib that will give the user object an API in which to interact with the filesystem.
#### The User
This code can be found in `user.go`.<br>
The user object will support functionality for creating a user instance and keeping track of special permissions associated with this uer.<br>

#### The Library
This code is in the file `lib.go`.<br>
The library object gives the user access to different utilities for interacting with files in virtual space; examples of these include versions of native Unix-like utilities such as `open`, 'close`, 'remove`, `ls` and so on.<br><br>
An extension of this object is the function `shellLoop()` which can be found in `shell.go`. This function provides the shell that the user will uuse to interact with the filesystem.

#### The Filesystem
See `filesystem.go`<br>
The collection of structures that will hold all the file and directories we'll be interacting with.

#### Onward to Ex02
When I've completed the next part, I'll post the link here (or you could just visit 02/, whatever works I guess ':D).











