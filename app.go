package main

// our main task runner which runs the app
func main() {
  inventory := new(Inventory).buildInventory()
  inventory.stockCheck()
}
