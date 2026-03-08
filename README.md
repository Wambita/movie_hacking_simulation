# Hacker CLI Simulator

![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![CLI](https://img.shields.io/badge/interface-CLI-blue)
![Status](https://img.shields.io/badge/project-simulation-orange)

A **movie-style hacker terminal simulator written in Go** that recreates the dramatic command-line scenes often seen in cybersecurity movies.

The program simulates:

* Terminal typing effects
* Matrix-style terminal noise
* Network scanning
* System resource monitoring
* Password cracking attempts
* AI intrusion detection training
* Progress bars
* Random warnings and errors

This project was built to **practice Go CLI development while creating a fun terminal simulation.**

---

# Demo

Example terminal output:

```text
INITIALIZING DARKNET TERMINAL...
LOADING EXPLOIT FRAMEWORK...

k29fj29f8fj392jf9f9jf2jf9
aj29fj2f9j2jf9j2f9j2fj2

Injecting firewall bypass module
[████████████████████████████] COMPLETE

Scanning target network...
Host detected -> 192.168.10.22 : ports [22,80,443]

Monitoring system resources...
CPU: 57% | MEMORY: 63% | THREADS: 2910

Running password cracking module...
Trying password: admin123 FAILED
Trying password: root FAILED
Password found: r00t@access ✔

ACCESS GRANTED ✔
```

---

# Features

##  Hacker Typing Effect

Text appears **character by character** like a real terminal being typed.

##  Matrix-Style Noise

Random characters simulate encrypted network traffic and terminal activity.

##  Network Scanner Simulation

Random IP addresses simulate discovered hosts.

##  System Resource Monitor

Displays simulated:

* CPU usage
* Memory usage
* Thread count

##  Password Cracking Module

Simulates brute-force password attempts.

##  AI Intrusion Detection Training

Simulates training logs with:

* epochs
* loss values
* accuracy metrics

##  Warning and Error Simulation

Random warnings and errors make the terminal feel realistic.

---

# Project Structure

```
hacker-cli-simulator
│
├── main.go
└── README.md
```

---

# Installation

Clone the repository:

```bash
git clone https://github.com/Wambita/hacker-cli-simulator.git
cd hacker-cli-simulator
```

---

# Run the Program

```bash
go run main.go
```

For the best experience, run it in a **full screen terminal**.

---

# Learning Goals

This project helped practice:

* Go CLI development
* ANSI terminal color formatting
* Time-based animations
* Random data simulation
* Structuring small terminal tools

---

# Future Improvements

Possible improvements:

* Interactive commands (`scan`, `breach`, `deploy`)
* Matrix-style falling characters
* Real keyboard input
* Terminal UI using `bubbletea`
* Multi-threaded simulations

---

# Author

**Sheila Fana**

GitHub:
- [https://github.com/Wambita](https://github.com/Wambita)

---

# License

MIT License

---
