package main

import (
	"fmt"
)

func main() {

	var member = map[int]string{
		321211: "Hendy",
		121211: "Alexia",
		131311: "sukemo",
	}

	// referance type
	// var newMember = member
	// newMember[121211] = "siapa"

	fmt.Println(member)
	// fmt.Println(newMember)

	for id, name := range member {
		fmt.Println(id, name)
	}

}
