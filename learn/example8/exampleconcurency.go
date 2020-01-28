package main

import (
	"fmt"
	"time"
)

func printConcurency(text string) {
	// fmt.Println("ini contoh concurency")
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(text)
	}
}

func main() {
	// go printConcurency()
	// time.Sleep(100 * time.Millisecond)
	// fmt.Println("welcome")

	start := time.Now()
	go printConcurency("ini contoh concurency")
	printConcurency("welcome")
	fmt.Println(time.Since(start))
}
