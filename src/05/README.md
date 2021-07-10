# Adding a Text Editor
Adding a text editor to make it possible to edit files in the filesystem.

The text editor for the fileystem makes use of the Chrome browser and the Ace Javascript library that gives access to editor-like functionality.<br>
One of the hardest challenges to this step was making an editor work for the Windows platform. The original idea was to make use of the <a href="https://github.com/bediger4000/kilo-in-go">Kilo-in-go</a> Golang terminal text editor to add a Vim-like editor to the filesystem. Though this was unsuccessful because of UNIX-like system calls the editor used and my inability to find equivalent Windows syscalls.<br><br>
So this editor would run from the browser.<br><br>
<img src="https://github.com/AlysonBee/GoVirtualFilesystem/blob/master/assets/webeditor.png"  height="400" width="500" />


## Putting together a "website" text editor

The first step was to write a web app in Go that could be turned on and off from the editor with the command "open". This I found to be a really Go specific thing because the original plan was to have a Python Flask app that could run and then be killed with signals but this didn't work at all.<br><br>

### Opening and Closing the App

This detail is reasonably important as it's the most creative part of this web app approach. Opening and closing the browser required a bit of out-of-the-box thinking. The main goal was to keep it as user-friendly and immersive as possible. I aimed to implement the functionality of Ctrl+S to save content and closing the browser tab as a way of closing the editor while also disabling the ability to crash the browser from the terminal. These three implementations went as follows;

#### Opening a File
```
open [filename]
```
A small segment from `shell.go` that's responsible for opening a new editor session;
```
    ....
		if _, exists := tmp.files[segments[len(segments)-1]]; exists {
			editingFile = tmp.files[segments[len(segments)-1]]
			editor()
			//s.readFile(tmp.files[segments[len(segments)-1]].rootPath)

		}
    ...
```
The function `editor()` is responsible for spinning up a new web session with the contents of the global variable `editingFile` which would dump the target file's contents into the browser by dereferencing the `content` section of the object (`editingFile.content`').<br>
<br><br>

The first few lines of code after initializing a new server is to disable the `Ctrl+C` functionality to make crashing it from the terminla throw out an error message telling the user to close the browser tab.
```
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println("Interrupt cancelled. Close text editor tab at :127.0.0.1:5000;", sig)
		}
	}()
```
The next piece of code opens the browser at the address `127.0.0.1:5000` and starts up our server to listen in on the same address.
```
  ...
	openbrowser("http://127.0.0.1:5000")
	server := http.Server{Addr: ":5000", Handler: mux}
  ...
```
Finally, our endpoint is mapped to the function `indexHandler`.
```
  mux.HandleFunc("/", indexHandler)
```
A global variable that sets up our template will allow us to send our `editingFile.content` to the HTML file `editor/editor.html`. The key issue with this design is the hardcoded path to `editor/editor.html`. With this path set this way, invoking the `open` command from anywhere that isn't the root of the directory will cause a crash. <b>This must be converted to an absolute path</b>.
 ```
 ...
 var tpl = template.Must(template.ParseFiles("editor/editor.html"))
 ...
 ```
 And lastly, our indexHandler code.
 ```
 	source := &sourceCode{
		Code: string(editingFile.content),
		Ext: "golang",
	}
	err := tpl.Execute(buf, source)
```
The Ace Editor library code responsible for displaying our source code in the editor inside `editor/editor.html`
```
...
    var editor = ace.edit("editor");
    editor.setTheme("ace/theme/monokai");
    editor.session.setMode("ace/mode/{{ .Ext }}");
    ...
    
    <div id="editor">{{ .Code }}</div>
    ...
...
```
#### Closing

Closing our text editor made use of Javascript Window interaction code. This made use of an AJAX trigger on the `beforeunload` window event (which basically means "before the window closes") which sent a call to the `/shutdown` endpoint in our Go code.<br><br>

Inside `editor/editor.html`

```
    $(window).on("beforeunload", function() { 
        $.ajax({
                type: 'POST',
                url: '/shutdown',
                contentType: 'application/json;charset=UTF-8',
                data: JSON.stringify({'data':'exiting'})
            })    
    })
```
Inside `editor.go`
```
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
    	if r.Method == "POST" {
    		server.Shutdown(context.Background())
    	}
    })
```
The above code shuts down the running browser session and returns control back to the filesystem.

#### Saving
Saving made use of Javascript code for key mapping code. When a keypress is detected and the simutaneous keys for `Ctrl` and `S` are pressed, the `save` endpoint is hit with a POST request to the backend.<br><br>

Inside `editor/editor.html`
```
    document.addEventListener("keydown", function(e) {
      if ((window.navigator.platform.match("Mac") ? e.metaKey : e.ctrlKey)  && e.keyCode == 83) {
        e.preventDefault();                         
        console.log(editor.getValue())
        $.ajax({
            type: 'POST',
            url: '/save',
            contentType: 'application/json;charset=UTF-8',
            data: JSON.stringify({'data': editor.getValue()}),
            success: function() {
                console.log("success")
            }
        });
   ```
Inside `editor.go`
   ```
   func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data map[string]string
		json.NewDecoder(r.Body).Decode(&data)
		editingFile.content = []byte(data["data"https://itnext.io/go-virtual-filesystem-adding-a-text-editor-176f082e0109])
	.}
}
```
The new "edited" code from the text editor overwrites the old value in `editingFile.content`.<br><br>

A more detailed version of this implementation can be seen in the corresponding article <a href="https://itnext.io/go-virtual-filesystem-adding-a-text-editor-176f082e0109">here</a>.






















