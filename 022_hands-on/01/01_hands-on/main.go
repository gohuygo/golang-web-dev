package main

import (
  "io"
  "net/http"
)

func main() {
  http.HandleFunc("/me", me)
  http.HandleFunc("/", root)
  http.HandleFunc("/dog", dog)
  http.ListenAndServe(":8080", nil)
}

func me(w http.ResponseWriter, r *http.Request) {
 io.WriteString(w, "Huy")
}

func dog(w http.ResponseWriter, r *http.Request) {
 io.WriteString(w, "dog")
}

func root(w http.ResponseWriter, r *http.Request) {
 io.WriteString(w, "root")
}
