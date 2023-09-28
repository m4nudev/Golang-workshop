package main

import "fmt"

func main() {
	// so here greeting is variable of type string
	var greeting string = "Hello World"
	fmt.Println(greeting)

	// wee can declare like that also
	var i int
	i = 45
	fmt.Println(i)

	// some Shorthand way to declare
	var s, t string = "foo", "bar"
	fmt.Println(s)
	fmt.Println(t)

	// different dataTypes can declare like this
	var (
		str string = "foo"
		x   int    = 5
	)
	fmt.Println(str)
	fmt.Println(x)

	// Short Variable Declaration
	//  := short variable assignment operator

	str1 := "Hello World"
	str1 = "namaste"
	fmt.Println(str1)
}
