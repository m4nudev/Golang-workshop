package main

import "fmt"

func main() {
	sum, difference := operation(20, 10)
	fmt.Println(sum, " ", difference)
}

func operation(a int, b int) (sum int, diff int) {
	sum = a + b
	diff = a - b
	return
}
