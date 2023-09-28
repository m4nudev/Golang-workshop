package main

import "fmt"

func main() {
	// Single input
	var name1 string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%s", &name1)
	fmt.Println("Hey there, ", name1)

	// Multiple input

	var name2 string
	var is_muggle bool

	fmt.Print("Enter your name & are you a muggle: ")
	fmt.Scanf("%s %t", &name2, &is_muggle)
	fmt.Println(name2, is_muggle)

	var a string
	var b int
	fmt.Print("Enter the string and a number: ")
	count, err := fmt.Scanf("%s %d", &a, &b)
	fmt.Println("count : ", count)
	fmt.Println("error : ", err)
	fmt.Println("a : ", a)
	fmt.Println("b : ", b)

}
