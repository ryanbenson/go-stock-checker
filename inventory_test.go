package main

import (
  "testing"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestInventory(t *testing.T) {
  g := Goblin(t)
  // hook in Gomega
  RegisterFailHandler(func(m string, _ ...int){ g.Fail(m) })

  g.Describe("#Inventory", func() {
    g.It("")
  })
}
