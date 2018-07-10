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
    g.It("")
  })
}
