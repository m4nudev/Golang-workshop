package main

import (
	"fmt"
	"time"
)

func main() {
	go sampleRoutine()
	fmt.Println("Started main")
	time.Sleep(1 * time.Second)
	fmt.Println("finished main")
}

func sampleRoutine() {
	fmt.Println("Inside sample Goroutine")
}

/*
output :
Started main
Inside sample Goroutine
finished main
*/
