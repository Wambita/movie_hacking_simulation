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

func matrixRain(lines int) {

	chars := "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*"

	for i := 0; i < lines; i++ {

		line := ""

		for j := 0; j < 70; j++ {
			line += string(chars[rand.Intn(len(chars))])
		}

		fmt.Println(line)

		time.Sleep(20 * time.Millisecond)
	}
}

func progressBar(task string) {

	fmt.Println(task)

	barLength := 30

	for i := 0; i <= barLength; i++ {

		bar := strings.Repeat("█", i) + strings.Repeat(" ", barLength-i)

		fmt.Printf("\r[%s]", bar)

		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println(" COMPLETE")
}
func randomIP() string {

	return fmt.Sprintf("%d.%d.%d.%d",
		rand.Intn(255),
		rand.Intn(255),
		rand.Intn(255),
		rand.Intn(255))
}
func scanNetwork() {

	typeLine("Scanning target network...", 20*time.Millisecond)

	for i := 0; i < 10; i++ {

		ip := randomIP()

		fmt.Printf("Host detected -> %s : ports [22, 80, 443]\n", ip)

		time.Sleep(120 * time.Millisecond)
	}
}
