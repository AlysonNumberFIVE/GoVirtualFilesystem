# Creating the First Feature

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
		Prompt: 			currentUser.username + "$>",
		HistoryFile:		"/tmp/readline.tmp",
		AutoComplete:		autoCompleter,
		InterruptPrompt:	"^C",
		EOFPrompt:			"exit",
	})
	if err != nil {
		log.Fatal(err)
	}
	return prompt
}
```
This function sets the prompt we'll be making use of in our shell. It's turned into a member of the user object because the only external variable it uses is the `username` variable that can be found in `user`.<br><br>
