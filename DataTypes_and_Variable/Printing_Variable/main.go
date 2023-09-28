package main

import "fmt"

func main() {
	// var name string = "leetcode"
	// var user string = "Ramu"

	// fmt.Print("Welcome to ", name, ", ", user)
	/*
		"print"
		Only problem with using print method is that is does not introduce a new linr after printing/
		Solution to this problem of nre line character \n , it's use to create new line.
	*/

	// fmt.Print(name, "\n")
	// fmt.Print(user)

	/*
		To make the printing of newline process even more easier to us. Its called "Println"
	*/

	// fmt.Println(name)
	// fmt.Println(user)

	/*
		format specifier "printf"
		%v - default format
		%T - type of the value
		%d - integers
		%c - character
		%q - quoted characters/strign
		%s - plain string
		%t - true/false
		%f - floating number
		%.2f - floating numbers upto 2 decimal places
	*/
	var name string = "john"
	var i int = 78

	fmt.Printf("Hey, %v! you have scored %d/100 in physics", name, i)

}
