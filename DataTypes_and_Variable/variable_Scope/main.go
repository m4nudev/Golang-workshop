package main

import "fmt"

/*
	# Inner blocks can access variables declared within outer blocks.
	# Outer blocks cannot access variables declared within inner blocks

	Local variables
	#	Declared inside a function or a block
	#	not accessible outside the function or a block
	#	can also be declared inside loop and conditional statement

	Global variables
	#	Declared outside a function or a block
	#	available throughout the lifetime of the program
	#	declared at the top of the program outside all functions or blocks
	#	can be accessed from any part of thr program
*/

var name string = "steve smith" // Global variable

func main() {
	city := "London" // local variable
	{
		country := "UK"
		fmt.Println(country)
		fmt.Println(city)
	}
	fmt.Println(city)
	fmt.Println(name)
}

/*
Zero values
int - 0
float64 - 0.0
string - ""
bool - false
pointers, functions, interfaces, maps - nil
*/
