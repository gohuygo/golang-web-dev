      package main

      import (
        "log"
        "net/http"
        "text/template"
      )

      // 1. Take the previous program and change it so that:
      // * func main uses http.Handle instead of http.HandleFunc

      // Contstraint: Do not change anything outside of func main

      // Hints:

      // [http.HandlerFunc](https://godoc.org/net/http#HandlerFunc)
      // ``` Go
      // type HandlerFunc func(ResponseWriter, *Request)
      // ```

      // [http.HandleFunc](https://godoc.org/net/http#HandleFunc)
      // ``` Go
      // func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
      // ```

      // [source code for HandleFunc](https://golang.org/src/net/http/server.go#L2069)
      // ``` Go
      //   func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
      //       mux.Handle(pattern, HandlerFunc(handler))
      //   }
      // ```

      var tpl *template.Template

      func init() {
        tpl = template.Must(template.ParseGlob("*.gohtml"))
      }

      func main() {
        http.Handle("/me", http.HandlerFunc(me))
        http.Handle("/", http.HandlerFunc(root))
        http.Handle("/dog", http.HandlerFunc(dog))
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
