package main

import (
  "log"
  "os"
  "text/template"
)

type hotel struct {
  Name string
  Address string
  City string
  Zip string
}

type region struct {
  Name string
  Hotels []hotel
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("hotel.gohtml"))
}

func main() {
  regions := []region{
    region{
      Name: "Southern",
      Hotels: []hotel{
        hotel{"Hotel Southern 1", "123 Test Street", "Los Angeles", "12345"},
        hotel{"Hotel Southern 2", "456 Test Street", "Los Angeles", "12345"},
        hotel{"Hotel Southern 3", "678 Test Street", "Los Angeles", "12345"},
      },
    },
    region{
      Name: "Central",
      Hotels: []hotel{
        hotel{"Hotel Central 1", "123 Test Street", "Los Angeles", "12345"},
        hotel{"Hotel Central 2", "456 Test Street", "Los Angeles", "12345"},
        hotel{"Hotel Central 3", "678 Test Street", "Los Angeles", "12345"},
      },
    },
    region{
      Name: "Northern",
      Hotels: []hotel{
        hotel{"Hotel Northern 1", "123 Test Street", "Los Angeles", "12345"},
        hotel{"Hotel Northern 2", "456 Test Street", "Los Angeles", "12345"},
        hotel{"Hotel Northern 3", "678 Test Street", "Los Angeles", "12345"},
      },
    },
  }

  err := tpl.Execute(os.Stdout, regions)
  if err != nil {
    log.Fatalln(err)
  }
}
