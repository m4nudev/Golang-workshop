package main

import "fmt"

func main() {
	factOf5 := fact(5)
	fmt.Println(factOf5)
}

func fact(n int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res *= i
	}
	return res
}

/*
output :
120
*/
