package main

import (
  "log"
  "os"
  "fmt"
  "text/template"
  "encoding/csv"
  "strconv"
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

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("stocks.gohtml"))
}

func main() {
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
      fmt.Println(open)

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

  tplErr := tpl.Execute(os.Stdout, stocks)
  
  if tplErr != nil {
    log.Fatalln(err)
  }
}
