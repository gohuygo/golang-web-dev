// # Building upon the code from the previous problem:

// Add code to respond to the following METHODS & ROUTES:
//   GET /
//   GET /apply
//   POST /apply

package main

import (
  "net"
  "log"
  "io"
  "bufio"
  "fmt"
  "strings"
)

func main() {
  li, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalln(err)
  }
  defer li.Close()

  for {
    conn, err := li.Accept()
    if err != nil {
      log.Println(err)
      continue
    }
    go serve(conn)
  }
}

func serve(conn net.Conn) {
  defer conn.Close()
  scanner := bufio.NewScanner(conn)
  i := 0
  var method, uri string

  for scanner.Scan() {
    ln := scanner.Text()
    if i == 0 {
      xs := strings.Fields(ln)
      method = xs[0]
      uri = xs[1]
      fmt.Println("METHOD:", method)
      fmt.Println("URI:", uri)
    }

    if ln == "" {
      fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
      break
    }

    i++
  }

  switch {
  case method == "GET" && uri == "/":
    handleIndex(conn)
  case method == "GET" && uri == "/apply":
    handleApply(conn)
  case method == "POST" && uri == "/apply":
    handleApplyPost(conn)
  default:
    handleDefault(conn)
  }

}

func handleIndex(c net.Conn) {
  body := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>GET INDEX</title>
    </head>
    <body>
      <h1>"GET INDEX"</h1>
      <a href="/">index</a><br>
      <a href="/apply">apply</a><br>
    </body>
    </html>
  `
  io.WriteString(c, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(c, "Content-Type: text/html\r\n")
  io.WriteString(c, "\r\n")
  io.WriteString(c, body)
}

func handleApply(c net.Conn) {
  body := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>GET DOG</title>
    </head>
    <body>
      <h1>"GET APPLY"</h1>
      <a href="/">index</a><br>
      <a href="/apply">apply</a><br>
      <form action="/apply" method="POST">
      <input type="hidden" value="In my good death">
      <input type="submit" value="submit">
      </form>
    </body>
    </html>
  `
  io.WriteString(c, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(c, "Content-Type: text/html\r\n")
  io.WriteString(c, "\r\n")
  io.WriteString(c, body)
}

func handleApplyPost(c net.Conn) {
  body := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>POST APPLY</title>
    </head>
    <body>
      <h1>"POST APPLY"</h1>
      <a href="/">index</a><br>
      <a href="/apply">apply</a><br>
    </body>
  </html>
  `
  io.WriteString(c, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(c, "Content-Type: text/html\r\n")
  io.WriteString(c, "\r\n")
  io.WriteString(c, body)
}

func handleDefault(c net.Conn) {
  body := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>default</title>
    </head>
    <body>
      <h1>"default"</h1>
    </body>
    </html>
  `
  io.WriteString(c, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(c, "Content-Type: text/html\r\n")
  io.WriteString(c, "\r\n")
  io.WriteString(c, body)
}

