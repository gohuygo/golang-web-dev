// # Building upon the code from the previous problem:

// Print to standard out (the terminal) the REQUEST method and the REQUEST URI from the REQUEST LINE.

// Add this data to your REPONSE so that this data is displayed in the browser.
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
 
  body := "METHOD IS " + method + "\n" + "URI IS " + uri

  io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
  fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
  fmt.Fprint(conn, "Content-Type: text/plain\r\n")
  io.WriteString(conn, "\r\n")
  io.WriteString(conn, body)


}

