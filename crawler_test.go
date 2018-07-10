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

    g.It("when getting a valid page", func() {
      doc, err := c.getPage("http://www.ryanbensonmedia.com")
      Ω(err).ShouldNot(HaveOccurred())
      Ω(doc).ShouldNot(BeNil())
    })

    g.It("when getting a valid website, but invalid page", func() {
      doc, err := c.getPage("http://warrendouglas.com/404")
      Ω(err).Should(HaveOccurred())
      Ω(doc).Should(BeNil())
    })

    g.It("when getting an invalid page", func() {
      doc, err := c.getPage("http://www.thiswillprobablynotexistihopethatwouldbesad.com")
      Ω(err).Should(HaveOccurred())
      Ω(doc).Should(BeNil())
    })
  })
}
