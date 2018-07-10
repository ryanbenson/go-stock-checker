package main

import (
  "os"
  "fmt"
  "log"
  "strings"
  "net/url"
  "net/http"

  "github.com/gen2brain/beeep"
)

type Notification struct{}

// notifySoldOut will send a notification when the item is still not available
// @params: nil
// @return: nil
func (n Notification) notifySoldOut(item Item) {
  fmt.Printf("%s : %s\n", item.name, item.soldOut)
}

// notifySoldOut will send notifications (text, STDOUT) when the item is available
// @params: nil
// @return: nil
func (n Notification) notifyInStock(item Item) {
  fmt.Printf("%s : %s\n", item.name, "IN STOCK")
  err := beeep.Notify("In Stock Item", item.url, "assets/warehouse.png")
  if err != nil {
      panic(err)
  }
  message := n.buildTextMessage(item)
  n.sendTextMessage(message)
}

// sendTextMessage will simply fire off a request to Twilio to send a text message
// @params: messageText (string)
// @return: nil
func (n Notification) sendTextMessage(messageText string) {
  url := "https://api.twilio.com/2010-04-01/Accounts/" + os.Getenv("TWILIO_SID") + "/Messages.json"

  msgData := n.getTextMessageData(messageText)
  msgDataReader := *strings.NewReader(msgData.Encode())

  req := n.buildTextMessageRequest(url, &msgDataReader)
  client := &http.Client{}
  resp, _ := client.Do(req)
  if (resp.StatusCode >= 200 && resp.StatusCode < 300) {
    log.Println("Text message sent")
  } else {
    log.Println(resp.Status);
  }
}

// getTextMessageData will build our API call message data set
// @params: messageText (string)
// @return: url.Values
func (n Notification) getTextMessageData(messageText string) url.Values {
  msgData := url.Values{}
  msgData.Set("To",os.Getenv("TWILIO_PHONE_NUMBER_TO"))
  msgData.Set("From",os.Getenv("TWILIO_PHONE_NUMBER_FROM"))
  msgData.Set("Body",messageText)
  return msgData
}

// buildTextMessageRequest builds our request to fire off with a client later
// @params: url (string)
// @params: msgDataReader (*strings.Reader)
// @return: *http.Request
func (n Notification) buildTextMessageRequest(url string, msgDataReader *strings.Reader) *http.Request {
  req, _ := http.NewRequest("POST", url, msgDataReader)
  req.SetBasicAuth(os.Getenv("TWILIO_SID"), os.Getenv("TWILIO_AUTH_TOKEN"))
  req.Header.Add("Accept", "application/json")
  req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
  return req
}

// buildTextMessage creates our text message that gets sent
// @params: item (Item)
// @return: string
func (n Notification) buildTextMessage(item Item) string {
  message := "ITEM IN STOCK: " + item.name + "\nClick the link to buy it now:\n" + item.url
  return message
}
