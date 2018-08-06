package main

import (
  "log"
  "net/http"
  "strconv"
)

func main() {
  http.HandleFunc("/", set)
  http.Handle("/favicon.ico", http.NotFoundHandler())
  http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
  var count int = 1

  cookie, err := req.Cookie("count")
  
  if err != nil {
    log.Println(err)
  } else {
    cookieValue, _ := strconv.Atoi(cookie.Value)
    count = cookieValue + 1
  }


  http.SetCookie(w, &http.Cookie{
    Name:  "count",
    Value: strconv.Itoa(count),
  })

}

