# Ex 02 - Creating the First Feature

This is the repo for my Medium post that can be found <a href="">here</a>.<br>
In this exercise, we'll be creating our first new features and refactoring our code a bit.<br>

## The Feature

The features to be added will be as follows.</br>
- The functionality to create a username on startup that will appear in our prompt in the shell.
- Command line features for autocomplete and command history.

<br>We'll be aiming to have a shell that has the following functionality:<br><br>
<img src="https://github.com/AlysonBee/GoVirtualFilesystem/blob/master/assets/demo1.gif"  height="400" />
## Inserting Username on Startup

Open the file `init.go`<br>
Inside `initUser`, you'll find the `setName()` function. All it will do for now is either take in a username of your choice or the input value `1` for potential 
guest users. Once done, this information will be used as an argument for our `createUser()` function to instantiate a new user with our chosen name.

## Passing our name over to the prompt.
Inside `user.go`<br>
The prompt makes use of the library `readline` (which can be downloaded and imported from `github.com/chzyer/readline` and used as shown in the imports in this file.
 ```
 import(
  ...
    "github.com/chzyer/readline
  )
  ```
  Here we see the first of our refactors. All of `readline`s functionality is added directly into the main body of the method `initPrompt()`.<br>
Let's go into this function a little bit more.
```
func (currentUser * user) initPrompt() (*readline.Instance) {
	autoCompleter := readline.NewPrefixCompleter(
		readline.PcItem("open"),
		readline.PcItem("close"),
		readline.PcItem("mkdir"),
		readline.PcItem("cd"),
		readline.PcItem("rmdir"),
		readline.PcItem("rm"),
		readline.PcItem("exit"),
	)
	prompt, err := readline.NewEx(&readline.Config{
		Prompt: currentUser.username + "$>",
		HistoryFile: "/tmp/readline.tmp",
		AutoComplete: autoCompleter,
		InterruptPrompt: "^C",
		EOFPrompt: "exit",
	})
	if err != nil {
		log.Fatal(err)
	}
	return prompt
}
```
This function sets the prompt we'll be making use of in our shell. It's turned into a member of the user object because the only external variable it uses is the `username` variable that can be found in `user`.<br><br>
### The Readline functions
Here's a brief rundown of the above code:<br>
Inside `readline.Config`:<br>
- `Prompt` - The name that will appear on the user prompt in the shell.
- `HistoryFile` - The location of the file that will store all command history.
- `AutoComplete` - The list of strings that can be autocompleted in the shell (by pressing TAB).
- `InterruptPrompt` - The signal handle that can interrupt the shell.
- `EOFPrompt` - The message printed out on exit.<br>


`NewPrefixCompleter` - The function responsible for handling autocompletion.<br><br>
## The Shell Loop
Open file `shell.go`<BR>
In this file resides our main; first it will initialize the project by allowing us to set our name, then pass its control to `shellLoop()`.<br>
In our shell loop, our `readline` prompt object is passed over to the shell and our Filesystem is initialized.<br><br>
### Changes in our filesystem.
`lib.go` was deleted and all it's member methods were passed over into the shell which are used in the `filesystem.execute()` command inside the `ShellLoop's body, like so;
The new new loop looks like this:
```
filesystem := initFilesystem()
prompt := currentUser.initPrompt()
for {
	input, _ := prompt.Readline()
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		continue 
	}
	filesystem.execute(input)
}
```
And `execute()` has the following body.<br>
```
// execute runs the commands passed into it.
func (fs * fileSystem) execute(command string) {

	switch command {
	case "open":
		fs.open()
	case "close":	
		fs.close()
	case "ls":
		fs.listDir()
	case "rm":
		fs.removeFile()
		fs.removeDir() 
	case "cd":
		fs.chDir()
	case "exit":
		fs.tearDown()
		os.Exit(1)
	default:
		fmt.Println(command, ": Command not found")
	}
}
```
## Onwards to 03
That's all for this week. On to fleshing out the `initFilesystem()` method next.
