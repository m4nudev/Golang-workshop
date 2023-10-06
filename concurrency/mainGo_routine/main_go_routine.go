package main

import (
	"fmt"
	"time"
)

func main() {
	go start()
	time.Sleep(1 * time.Second)
}

func start() {
	go process()
	fmt.Println("In start")
}

func process() {
	fmt.Println("In process")
}

// pather panchali, the black in white drama, garm hava.
// Scarface. irreversibke
