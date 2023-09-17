package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)
	factOfn := fact(n)
	fmt.Printf("factorial of %v is %v\n: ", n, factOfn)
}

func fact(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

/*
input : 6
output : 720
*/
