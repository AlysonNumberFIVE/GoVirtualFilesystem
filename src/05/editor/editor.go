
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"io/ioutil"
	"bytes"
)

var tpl = template.Must(template.ParseFiles("editor.html"))

type sourceCode struct {
	Code string
	Ext string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	buf := &bytes.Buffer{}
	content, _ := ioutil.ReadFile("TESTING")
	source := &sourceCode{
		Code: string(content),
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
		fmt.Println("POSTED")
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/save", saveHandler)
	http.ListenAndServe(":5001", mux)
}




