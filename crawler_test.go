package main

import (
  "testing"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestCrawler(t *testing.T) {
  g := Goblin(t)
  // hook in Gomega
  RegisterFailHandler(func(m string, _ ...int){ g.Fail(m) })

  g.Describe("#Crawler", func() {
    c := Crawler{}

    g.It("when getting a valid page, should receive the page", func() {
      page, err := c.getPage("http://warrendouglas.com")
      Ω(err).ShouldNot(HaveOccurred())
      Ω(page).ShouldNot(BeNil())
    })

    g.It("when getting a valid website, but invalid page, should receive an error", func() {
      page, err := c.getPage("http://warrendouglas.com/404")
      Ω(err).Should(HaveOccurred())
      Ω(page).Should(BeNil())
    })

    g.It("when getting an invalid page, but cannot reach, should receive an error", func() {
      page, err := c.getPage("http://www.thiswillprobablynotexistihopethatwouldbesad.com")
      Ω(err).Should(HaveOccurred())
      Ω(page).Should(BeNil())
    })

    g.It("when parsing HTML to get the DOM, should receive the DOM", func() {
      page, _ := c.getPage("http://warrendouglas.com")
      doc, err := c.getHTMLDoc(page)
      Ω(err).ShouldNot(HaveOccurred())
      Ω(doc).ShouldNot(BeNil())
    })

    g.It("when getting the text from a node, should receive a string", func() {
      page, _ := c.getPage("http://warrendouglas.com/about")
      doc, _ := c.getHTMLDoc(page)
      text := c.getNodeText(doc, ".about-content h2")
      Ω(text).Should(Equal("Welcome. We are The WD."))
    })
  })
}
