package main

import (
  "fmt"
  "log"
  "strings"
  "net/http"

  "github.com/PuerkitoBio/goquery"
  "github.com/gen2brain/beeep"
)

func ExampleScrape() {
  // Request the HTML page.
  res, err := http.Get("https://www.treesofantiquity.com/index.php?main_page=product_info&products_id=53")
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
      fmt.Println("https://www.treesofantiquity.com/index.php?main_page=product_info&products_id=53 : Sold Out")
    } else {
      err := beeep.Notify("In Stock Item", "https://www.treesofantiquity.com/index.php?main_page=product_info&products_id=53", "assets/warehouse.png")
      if err != nil {
          panic(err)
      }
    }
  })
}

func main() {
  ExampleScrape()
}
