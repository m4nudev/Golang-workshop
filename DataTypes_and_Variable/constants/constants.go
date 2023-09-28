package main

import "fmt"

const PI float64 = 3.14 // global variable

func main() {
	//-------------untyped constant----------------
	const name = "Harry Potter"
	const is_muggle = true
	const age = 24

	fmt.Printf("%v: %T \n", name, name)
	fmt.Printf("%v: %T \n", is_muggle, is_muggle)
	fmt.Printf("%v: %T \n", age, age)

	//-------------typed constant---------------- sue case of constant

	var radius float64 = 5.0
	var area float64
	area = PI * radius * radius
	fmt.Println("Area of the circle is : ", area)

}
