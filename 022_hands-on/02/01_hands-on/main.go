// # Building upon the code from the previous problem:

// Change your RESPONSE HEADER "content-type" from "text/plain" to "text/html"

// Change the RESPONSE from "CHECK OUT THE RESPONSE BODY PAYLOAD" (and everything else it contained: request method, request URI) to an HTML PAGE that prints "HOLY COW THIS IS LOW LEVEL" in <h1> tags.
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
 
  body := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
      <meta charset="UTF-8">
      <title>TEST</title>
    </head>
    <body>
      <h1>"HOLY COW THIS IS LOW LEVEL"</h1>
    </body>
    </html>
  `
  io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(conn, "Content-Type: text/html\r\n")
  io.WriteString(conn, "\r\n")
  io.WriteString(conn, body)


}

