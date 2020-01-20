package main

import (
	"fmt"
)

type Member struct {
	name string
	age  int
}

func (m Member) getInfo() {
	fmt.Println("hai saya di panggil ini infonya")
}

func (c *Member) changeName(newName string, newAge int) {
	c.name = newName
	c.age = newAge
}

func main() {
	user := Member{"Hendy", 30}
	user.getInfo()
	user.changeName("Hero", 21)
	fmt.Println()
}
