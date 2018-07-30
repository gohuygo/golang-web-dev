package main

import (
  "log"
  "net/http"
  "text/template"
)


// 1. Take the previous program in the previous folder and change it so that:
// * a template is parsed and served
// * you pass data into the template

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
  http.HandleFunc("/me", me)
  http.HandleFunc("/", root)
  http.HandleFunc("/dog", dog)
  http.ListenAndServe(":8080", nil)
}

func me(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "index.gohtml", "Huy")
  if err != nil {
    log.Fatalln(err)
  }
}

func dog(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "index.gohtml", "dog")
  if err != nil {
    log.Fatalln(err)
  }
}

func root(w http.ResponseWriter, r *http.Request) {
  err := tpl.ExecuteTemplate(w, "index.gohtml", "root")
  if err != nil {
    log.Fatalln(err)
  }
}
