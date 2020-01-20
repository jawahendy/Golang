package main

import (
	"fmt"
)

func main() {
	var member []string

	member[0] = "Hendy"
	member[1] = "dol"
	member[2] = "hendybung"

	angka := [...]int{10, 100, 20, 30}
	angka[1] = 121

	newMember := make([]string, 3)
	copy(newMember, member)

	newMember[1] = "joko"
	newMember = append(newMember, "joki", "mukena")

	for i := 0; i < len(member); i++ {
		fmt.Println("Hello nama saya " + member[i])
	}

	for index, value := range member {
		fmt.Println("member %d = %s\n", index, value)
	}

	// sliceMember := member[1:2]
	sliceMember := member[:2]
	// sliceMember[1] = "Joko"

	fmt.Println(newMember)

	fmt.Println(sliceMember)
	fmt.Println("Length ", len(sliceMember))
	fmt.Println("Capacity ", cap(sliceMember))
	// fmt.Println(member[0])
	// fmt.Println(member[1])
	fmt.Println(angka)
}
