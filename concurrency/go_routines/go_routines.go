package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 1; i <= 10000; i++ {
		go calculateSquare(i)
	}
	elapsed := time.Since(start)
	time.Sleep(2 * time.Second)
	fmt.Println("Function took ", elapsed)
}

func calculateSquare(i int) {
	time.Sleep(1 * time.Second)
	fmt.Println(i * 1)
}
