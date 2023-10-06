package main

import "fmt"

func main() {
	v, _ := f()
	fmt.Println(v)
}

func f() (int, int) {
	return 42, 53
}
