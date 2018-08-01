// # Starting with the code in the "starting-files" folder:

// ## wire this program up so that it works

// ParseGlob in an init function

// Use HandleFunc for each of the routes

// Combine apply & applyProcess into one func called "apply"

// Inside the func "apply", use this code to create the logic to respond differently to a POST method and a GET method

// ``` go
// if req.Method == http.MethodPost {
//     		// code here
//     		return
//     	}
// ```

package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template


func init() {
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/apply", apply)
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		applyProcess(w, req)
		return
	}

	err := tpl.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, req *http.Request) {
	err := tpl.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
