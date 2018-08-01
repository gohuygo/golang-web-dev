package main

import(
  "net/http"
  "log"
  "html/template"
)

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
  http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
  http.HandleFunc("/", dogs)

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func dogs(w http.ResponseWriter, req *http.Request) {
  err := tpl.Execute(w, nil)
  if err != nil {
    log.Fatalln("template didn't execute: ", err)
  }
}

