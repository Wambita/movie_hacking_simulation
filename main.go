package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	green  = "\033[32m"
	red    = "\033[31m"
	yellow = "\033[33m"
	cyan   = "\033[36m"
	reset  = "\033[0m"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func typeLine(text string, speed time.Duration) {
	for _, c := range text {
		fmt.Printf("%c", c)
		time.Sleep(speed)
	}
	fmt.Println()
}

func matrixRain(lines int) {

	chars := "abcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*"

	for i := 0; i < lines; i++ {

		line := ""

		for j := 0; j < 70; j++ {
			line += string(chars[r.Intn(len(chars))])
		}

		fmt.Println(line)

		time.Sleep(20 * time.Millisecond)
	}
}

func randomIP() string {

	return fmt.Sprintf("%d.%d.%d.%d",
		r.Intn(255),
		r.Intn(255),
		r.Intn(255),
		r.Intn(255))
}

func randomWarning() {

	if r.Intn(5) == 1 {

		warnings := []string{
			"Packet loss detected...",
			"Connection unstable...",
			"Firewall resistance detected...",
			"High latency node detected...",
		}

		fmt.Println(yellow + "WARNING: " + warnings[r.Intn(len(warnings))] + reset)
	}
}

func randomError() {

	if r.Intn(8) == 2 {

		errors := []string{
			"Permission denied.",
			"Exploit attempt failed.",
			"Timeout connecting to node.",
		}

		fmt.Println(red + "ERROR: " + errors[r.Intn(len(errors))] + reset)
	}
}

func progressBar(task string) {

	fmt.Println(cyan + task + reset)

	barLength := 30

	for i := 0; i <= barLength; i++ {

		bar := strings.Repeat("█", i) + strings.Repeat(" ", barLength-i)

		fmt.Printf("\r[%s]", bar)

		time.Sleep(time.Duration(r.Intn(70)) * time.Millisecond)
	}

	fmt.Println(" COMPLETE")
}

func scanNetwork() {

	typeLine("Scanning target network...", 20*time.Millisecond)

	for i := 0; i < 10; i++ {

		ip := randomIP()

		fmt.Printf("Host detected -> %s : ports [22, 80, 443]\n", ip)

		randomWarning()
		randomError()

		time.Sleep(120 * time.Millisecond)
	}
}

func cpuSimulation() {

	fmt.Println(cyan + "Monitoring system resources..." + reset)

	for i := 0; i < 10; i++ {

		cpu := r.Intn(60) + 30
		mem := r.Intn(40) + 40

		fmt.Printf("CPU: %d%% | MEMORY: %d%% | THREADS: %d\n",
			cpu,
			mem,
			r.Intn(4000))

		time.Sleep(300 * time.Millisecond)
	}
}

func passwordCrack() {

	fmt.Println(cyan + "Running password cracking module..." + reset)

	passwords := []string{
		"admin123",
		"root",
		"qwerty",
		"letmein",
		"shadow",
		"dragon",
	}

	for i := 0; i < 12; i++ {

		pass := passwords[r.Intn(len(passwords))]

		fmt.Printf("Trying password: %-10s ", pass)

		time.Sleep(120 * time.Millisecond)

		fmt.Println(red + "FAILED" + reset)
	}

	fmt.Println(green + "Password found: r00t@access ✔" + reset)
}

func aiTrainingSimulation() {

	fmt.Println(cyan + "Launching AI intrusion detection model..." + reset)

	for i := 1; i <= 10; i++ {

		loss := r.Float64()

		fmt.Printf("Epoch %d/10 - loss: %.4f - accuracy: %.2f%%\n",
			i,
			loss,
			r.Float64()*100)

		time.Sleep(350 * time.Millisecond)
	}
}

func main() {

	fmt.Print(green)

	typeLine("INITIALIZING DARKNET TERMINAL...", 25*time.Millisecond)

	typeLine("LOADING EXPLOIT FRAMEWORK...", 25*time.Millisecond)

	time.Sleep(time.Second)

	matrixRain(20)

	progressBar("Injecting firewall bypass module")

	scanNetwork()

	cpuSimulation()

	progressBar("Deploying packet sniffer")

	typeLine("Executing automated exploit toolkit...", 25*time.Millisecond)

	typeLine("Please wait while the tool runs...", 25*time.Millisecond)

	matrixRain(15)

	passwordCrack()

	aiTrainingSimulation()

	progressBar("Extracting encrypted payload")

	fmt.Println(green + "ACCESS GRANTED ✔" + reset)

	fmt.Println("Session terminated.")
}
