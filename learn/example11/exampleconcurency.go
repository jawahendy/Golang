package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// wg have 3  Add , Wait and Done

	wg.Add(2)
	go printText("salam", &wg)
	go printText("sejahtera", &wg)

	wg.Wait()
}

func printText(text string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
	}
	wg.Done()
}
