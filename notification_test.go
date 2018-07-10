package main

import (
  "testing"
  . "github.com/franela/goblin"
  . "github.com/onsi/gomega"
)

func TestNotification(t *testing.T) {
  g := Goblin(t)
  // hook in Gomega
  RegisterFailHandler(func(m string, _ ...int){ g.Fail(m) })

  g.Describe("#Notification", func() {
    g.It("")
  })
}
