package main

import (
  "log"
  "os"
  "text/template"
)

type restaurant struct {
  Name string
  Menus []menu
}

type item struct {
  Name string
  Price float64
}

type menu struct {
  Name string //Breakfast, Lunch, Dinner
  Items []item
}


var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("menu.gohtml"))
}

func main() {
  restaurants := []restaurant{
    restaurant{
      Name: "Tag",
      Menus: []menu{
        {
          Name: "Breakfast",
          Items: []item{
            item{"Eggs", 1.00},
            item{"Fruits", 1.25},
            item{"Meat", 1.50},
          },
        },
        {
          Name: "Lunch",
          Items: []item{
            item{"Sandwich", 5.00},
            item{"Hotdog", 5.25},
            item{"Salad", 5.50},
          },
        },
        {
          Name: "Dinner",
          Items: []item{
            item{"Pizza", 10.00},
            item{"BBQ", 10.25},
            item{"Corndog", 10.50},
          },
        },
      },
    },

    restaurant{
      Name: "Second Restaurant Name",
      Menus: []menu{
        {
          Name: "Breakfast",
          Items: []item{
            item{"Eggs", 1.00},
            item{"Fruits", 1.25},
            item{"Meat", 1.50},
          },
        },
        {
          Name: "Lunch",
          Items: []item{
            item{"Sandwich", 5.00},
            item{"Hotdog", 5.25},
            item{"Salad", 5.50},
          },
        },
        {
          Name: "Dinner",
          Items: []item{
            item{"Pizza", 10.00},
            item{"BBQ", 10.25},
            item{"Corndog", 10.50},
          },
        },
      },
    },
  }

  err := tpl.Execute(os.Stdout, restaurants)
  if err != nil {
    log.Fatalln(err)
  }
}
