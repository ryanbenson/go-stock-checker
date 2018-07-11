package main

import (
  "testing"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestItem(t *testing.T) {
  g := Goblin(t)
  // hook in Gomega
  RegisterFailHandler(func(m string, _ ...int){ g.Fail(m) })

  g.Describe("#Item", func() {
    g.It("when checking for sold out status, it returns true if the text is matched", func() {
      i := Item{name: "Apple", url: "http://apples.com", dom: "#soldout", soldOut: "Sold Out"}
      websiteText := "Sold Out"
      soldOut := i.isSoldOut(websiteText)
      Ω(soldOut).Should(BeTrue())
    })

    g.It("when checking for sold out status, it returns false if the text is not matched", func() {
      i := Item{name: "Apple", url: "http://apples.com", dom: "#soldout", soldOut: "Sold Out"}
      websiteText := "Unavailable"
      soldOut := i.isSoldOut(websiteText)
      Ω(soldOut).Should(BeFalse())
    })

    g.It("when processing text, it removes all whitespace, and lowercased", func() {
      i := Item{name: "Apple", url: "http://apples.com", dom: "#soldout", soldOut: "\n\r Sold Out \n\r"}
      cleanSoldOut := i.cleanText(i.soldOut)
      Ω(cleanSoldOut).Should(Equal("sold out"))
    })
  })
}
