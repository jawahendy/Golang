package main

import (
	"fmt"
)

var member = map[int]string{
	321211: "Hendy",
	121211: "Alexia",
	131311: "sukemo",
}

func main() {

	// name, exists := member[121211]

	// fmt.Println(name)
	// fmt.Println(exists)

	// delete(member, 131311)

	if checkmember(121211) {
		fmt.Println("member dengan id ini dah ada")
	} else {
		fmt.Println("member gak ada")
	}

}

func checkmember(id int) bool {
	_, exists := member[id]
	return exists
}
