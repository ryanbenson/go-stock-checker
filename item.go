package main

import (
  "strings"
)

// Item to hold our item information
type Item struct {
  name string
  url string
  dom string
  soldOut string
}

// isSoldOut checks to see if the dom reference still says sold out or not
// @params: reference (string)
// @params: item (Item)
// @return: bool
func (i Item) isSoldOut(reference string) bool {
  var isSoldOut bool
  if(reference == i.soldOut) {
    isSoldOut = true
  } else {
    isSoldOut = false
  }
  return isSoldOut
}

// cleanText will standardize the given string to make it easier to compare by
// removing leading and trailing space, new lines, and make the string lowercase
// @params: str (string)
// @return: string
func (i Item) cleanText(str string) string {
  return strings.ToLower(strings.TrimSpace(str))
}
