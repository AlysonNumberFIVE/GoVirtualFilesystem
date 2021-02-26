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
