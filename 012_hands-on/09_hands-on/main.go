package main

import (
  "log"
  "os"
  "text/template"
  "encoding/csv"
  "strconv"
  "net/http"

)

type Stock struct {
  Date string
  Open float64
  High float64
  Low float64
  Close float64
  Volume int64
  AdjClose float64
}


func main() {
  http.HandleFunc("/parse", parse)
  http.ListenAndServe(":8080", nil)
}

func parse(res http.ResponseWriter, req *http.Request) {
  csvIn, errOpen := os.Open("table.csv")
 
  if errOpen != nil {
    log.Fatal(errOpen)
  }

  records, err := csv.NewReader(csvIn).ReadAll()
  if err != nil {
    log.Fatal(err)
  }
  var stocks []Stock
 
  for i, record := range records{

    if i > 0 && i <= 2 {
      open, err     := strconv.ParseFloat(record[1], 64)
      high, err     := strconv.ParseFloat(record[2], 64)
      low, err      := strconv.ParseFloat(record[3], 64)
      close, err    := strconv.ParseFloat(record[4], 64)
      volume, err   := strconv.ParseInt(record[5], 10, 64)
      adjClose, err := strconv.ParseFloat(record[6], 64)
     
      if err != nil {
        log.Fatal(err)
      }

      stocks = append(stocks, Stock{
        Date: record[0],
        Open: open,
        High: high,
        Low: low,
        Close: close,
        Volume: volume,
        AdjClose: adjClose,
      })
    }
  }
  
  tpl, err := template.ParseFiles("stocks.gohtml")
  err = tpl.Execute(res, stocks)
  
  if err != nil {
    log.Fatalln(err)
  }
}
