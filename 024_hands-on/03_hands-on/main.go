package main

// Use "http.FileServer"

import(
  "net/http"
  "log"
)


func main() {
  http.Handle("/", http.FileServer(http.Dir("./starting-files"))))


  log.Fatal(http.ListenAndServe(":8080", nil)
}
