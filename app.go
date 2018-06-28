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
  url string
  dom string
  soldOut string
}

func StockCheck() {
  inventory := Inventory{url: "https://www.treesofantiquity.com/index.php?main_page=product_info&products_id=53", dom: "#cartAdd", soldOut: "sold out"}
  // Request the HTML page.
  res, err := http.Get(inventory.url)
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
  doc.Find("#cartAdd").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    button_text := strings.ToLower(strings.TrimSpace(s.Find("span").Text()))
    if(button_text == "sold out") {
      fmt.Printf("%s : %s", inventory.url, inventory.soldOut)
    } else {
      err := beeep.Notify("In Stock Item", inventory.url, "assets/warehouse.png")
      if err != nil {
          panic(err)
      }
    }
  })
}

func main() {
  StockCheck()
}
