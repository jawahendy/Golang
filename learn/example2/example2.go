package main

import (
	"fmt"
)

func main() {
	hewan := [3][2]string{
		{"burung", "gagak"},
		{"sapi", "kambing"},
		{"puyuh", "godzila"},
	}

	for i := 0; i < len(hewan); i++ {
		for j := 0; j < len(hewan[i]); j++ {
			fmt.Println(hewan[i][j])
		}
	}

	fmt.Println(hewan)
	fmt.Println(hewan[2][1])
}
