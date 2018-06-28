package main

import (
  "io"
  "os"
  "fmt"
  "log"
  "strings"
  "net/url"
  "net/http"
  "encoding/json"

  "github.com/PuerkitoBio/goquery"
  "github.com/gen2brain/beeep"
  _ "github.com/joho/godotenv/autoload"
)

type Inventory struct {
  items []Item
}

type Item struct {
  name string
  url string
  dom string
  soldOut string
}

func StockCheck() {
  inventory := buildInventory()

  for _, item := range inventory.items {
    // Request the HTML page.
    body, err := getPage(item.url)
    if err != nil {
      log.Print(err)
      continue
    }

    doc, err := getHtmlDoc(body)
    if err != nil {
      log.Print(err)
      continue
    }

    // Find the review items
    doc.Find(item.dom).Each(func(i int, s *goquery.Selection) {
      // For each item found, get the band and title
      button_text := strings.ToLower(strings.TrimSpace(s.Text()))
      if(button_text == "sold out") {
        fmt.Printf("%s : %s\n", item.url, item.soldOut)
      } else {
        fmt.Printf("%s : %s\n", item.url, "IN STOCK")
        err := beeep.Notify("In Stock Item", item.url, "assets/warehouse.png")
        if err != nil {
            panic(err)
        }
        message := buildTextMessage(item)
        sendTextMessage(message)
      }
    })
  }
}

func getPage(url string) (io.ReadCloser, error) {
  res, err := http.Get(url)
  if err != nil {
    return nil, err
  }
  if res.StatusCode != 200 {
    return nil, err
  }
  return res.Body, nil
}

func getHtmlDoc(body io.ReadCloser) (*goquery.Document, error) {
  doc, err := goquery.NewDocumentFromReader(body)
  if err != nil {
    return nil, err
  }
  return doc, nil
}

func sendTextMessage(messageText string) {
  urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + os.Getenv("TWILIO_SID") + "/Messages.json"

  msgData := url.Values{}
  msgData.Set("To",os.Getenv("TWILIO_PHONE_NUMBER_TO"))
  msgData.Set("From",os.Getenv("TWILIO_PHONE_NUMBER_FROM"))
  msgData.Set("Body",messageText)
  msgDataReader := *strings.NewReader(msgData.Encode())

  client := &http.Client{}
  req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
  req.SetBasicAuth(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

  resp, _ := client.Do(req)
  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    var data map[string]interface{}
    decoder := json.NewDecoder(resp.Body)
    err := decoder.Decode(&data)
    if (err == nil) {
      fmt.Println(data["sid"])
    }
  } else {
    fmt.Println(resp.Status);
  }
}

func buildTextMessage(item Item) string {
  message := "ITEM IN STOCK: " + item.name + "\nClick the link to buy it now:\n" + item.url
  return message
}

func buildInventory() Inventory {
  allItems := []Item{
    Item{name: "Golden Russet", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&products_id=53", dom: "#cartAdd", soldOut: "sold out"},
    Item{name: "Black Oxford", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&cPath=1&products_id=639", dom: "#cartAdd", soldOut: "sold out"},
    Item{name: "Tree Starter Package", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&cPath=44&products_id=3", dom: "#cartAdd", soldOut: "sold out"},
  }
  inventory := Inventory{items: allItems}
  return inventory
}

func main() {
  StockCheck()
}
