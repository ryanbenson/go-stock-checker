package main

import (
  "fmt"
  "log"
  "strings"
  "net/http"

  "github.com/PuerkitoBio/goquery"
  "github.com/gen2brain/beeep"
)

type Inventory struct {
  name string
  url string
  dom string
  soldOut string
}

func StockCheck() {
  inventory := buildInventory()

  for _, item := range inventory {
    // Request the HTML page.
    res, err := http.Get(item.url)
    if err != nil {
      log.Fatal(err)
    }
    defer res.Body.Close()
    if res.StatusCode != 200 {
      log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
    }

    // Load the HTML document
    doc, err := goquery.NewDocumentFromReader(res.Body)
    if err != nil {
      log.Fatal(err)
    }

    // Find the review items
    doc.Find(item.dom).Each(func(i int, s *goquery.Selection) {
      // For each item found, get the band and title
      button_text := strings.ToLower(strings.TrimSpace(s.Find("span").Text()))
      if(button_text == "sold out") {
        fmt.Printf("%s : %s\n", item.url, item.soldOut)
      } else {
        fmt.Printf("%s : %s\n", item.url, "IN STOCK")
        err := beeep.Notify("In Stock Item", item.url, "assets/warehouse.png")
        if err != nil {
            panic(err)
        }
      }
    })
  }
}

func buildInventory() []Inventory {
  fullInventory := []Inventory{
    Inventory{name: "Golden Russet", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&products_id=53", dom: "#cartAdd", soldOut: "sold out"},
    Inventory{name: "Black Oxford", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&cPath=1&products_id=639", dom: "#cartAdd", soldOut: "sold out"},
    Inventory{name: "Tree Starter Package", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&cPath=44&products_id=3", dom: "#cartAdd", soldOut: "sold out"},
  }
  return fullInventory
}

func main() {
  StockCheck()
}
