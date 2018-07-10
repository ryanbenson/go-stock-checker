package main

import (
  "log"
)

// Inventory to hold all of our items
type Inventory struct {
  items []Item
}

// StockCheck checks out list of products to see if they are still sold out
// @params: nil
// @return: nil
func (i Inventory) stockCheck() {
  crawler := Crawler{}
  notification := Notification{}

  for _, it := range i.items {
    body, err := crawler.getPage(it.url)
    if err != nil {
      log.Print(err)
      continue
    }

    doc, err := crawler.getHtmlDoc(body)
    if err != nil {
      log.Print(err)
      continue
    }

    reference := crawler.getNodeText(doc, it.dom)

    isSoldOut := it.isSoldOut(reference)
    if isSoldOut == true {
      notification.notifySoldOut(it)
    } else {
      notification.notifyInStock(it)
    }
  }
}

// buildInventory creates the inventory of items to check
// params: nil
// @return: Inventory
func (i Inventory) buildInventory() Inventory {
  allItems := []Item{
    Item{name: "Golden Russet", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&products_id=53", dom: "#cartAdd", soldOut: "sold out"},
    Item{name: "Black Oxford", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&cPath=1&products_id=639", dom: "#cartAdd", soldOut: "sold out"},
    Item{name: "Tree Starter Package", url: "https://www.treesofantiquity.com/index.php?main_page=product_info&cPath=44&products_id=3", dom: "#cartAdd", soldOut: "sold out"},
  }
  inventory := Inventory{items: allItems}
  return inventory
}
