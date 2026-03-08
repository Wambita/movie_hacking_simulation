package main

import (
	"fmt"
	"time"
)
const (
	green = "\033[32m"
	reset = "\033[0m"
)


func typeLine(text string, speed time.Duration) {
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(speed)
	}
	fmt.Println()
}


func main() {
	fmt.Print(green)
	typeLine("INITIALIZING DARKNET TERMINAL...", 25*time.Millisecond)
}

