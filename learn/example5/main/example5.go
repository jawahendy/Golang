package main

import (
	"fmt"
	// "./learn/example5/models"
)

type Member struct {
	Name string
	age  int
}

func main() {

	user1 := Member{"hendy", 27}

	// user1 := models.Member{"hendy", 27}

	user2 := user1
	user2.Name = "Hendy aja"

	if user1 == user2 {
		fmt.Println("Betul mereka sama")
	} else {
		fmt.Println("maaf sudah beda cari yg lain")
	}

	fmt.Println(user2)
	fmt.Println("Hello nama saya" + user1.Name)
}
