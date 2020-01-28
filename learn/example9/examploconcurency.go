package main

import (
	"fmt"
	"time"
)

// channnel
var sentanceChannel = make(chan string)
var CleanSentanceChannel = make(chan string)

func main() {
	sentance := [7]string{"hendy", "belajar", "golang", "ngulang", "aja", "never", "give up"}

	go kalimat(sentance)
	go clean()
	go Save()

	time.Sleep(500 * time.Millisecond)
}

func kalimat(items [7]string) {
	for _, item := range items {
		if item == "hendy" {
			fmt.Println("berhasil dapat")
			sentanceChannel <- item
		}
	}
}

func clean() {
	for rawItem := range sentanceChannel {
		fmt.Println("berhasil clean" + rawItem)
		CleanSentanceChannel <- "harta clean"
	}
}

func Save() {
	for cleanedSentance := range CleanSentanceChannel {
		fmt.Println("berhasil simpan" + cleanedSentance)
	}
}
