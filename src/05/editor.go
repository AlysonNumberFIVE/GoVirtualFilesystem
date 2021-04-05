
package main

import (
	"fmt"
	"context"
	"html/template"
	"net/http"
	"bytes"
	"runtime"
	"os/signal"
	"os/exec"
	"os"
	"log"
	"encoding/json"
)

var tpl = template.Must(template.ParseFiles("editor/editor.html"))

type sourceCode struct {
	Code string
	Ext string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	
	source := &sourceCode{
		Code: string(editingFile.content),
		Ext: "golang",
	}
	err := tpl.Execute(buf, source)
	if err != nil {
		fmt.Println("Error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var data map[string]string
		json.NewDecoder(r.Body).Decode(&data)
		editingFile.content = []byte(data["data"])
	}
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}


func editor() {
	mux := http.NewServeMux()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println("Interrupt cancelled. Close text editor tab at :127.0.0.1:5000;", sig)
		}
	}()

	openbrowser("http://127.0.0.1:5000")
	server := http.Server{Addr: ":5000", Handler: mux}
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/save", saveHandler)
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
    	if r.Method == "POST" {
    		server.Shutdown(context.Background())
    	}
    })
	server.ListenAndServe()
}




