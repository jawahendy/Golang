package main

import "fmt"

func main() {
	var rows int = 5
	var k int

	for i := 1; i <= rows; i++ {
		k = 0
		for space := 1; space <= rows-i; space++ {
			fmt.Print("  ")
		}
		for {
			fmt.Print("* ")
			k++
			if k == 2*i-1 {
				break
			}
		}
		fmt.Println("")
	}

	// // untuk versi balik star
	// var rows = 5
	// var i, j int
	// for i = rows; i >= 1; i-- {
	// 	for space := 1; space <= rows-i; space++ {
	// 		fmt.Print("  ")
	// 	}
	// 	for j = i; j <= 2*i-1; j++ {
	// 		fmt.Printf("* ")
	// 	}
	// 	for j = 0; j < i-1; j++ {
	// 		fmt.Printf("* ")
	// 	}
	// 	fmt.Println("")
	// }
}
