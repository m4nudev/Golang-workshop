/*
Functions :
			- self contained units of code which carry out a certain job.
			- help us divide a program into a small manageable, repeatable and organisable chunks.

Why use functions?
			- Reusability
			- Abstraction

function syntax :
			- func <function_name> (<params>) <return type> {
				// body of the function
			}

			func addNumbers(a int, b int) int {
				..bosy of the function
			}

			Calling a function
			<function_name> (<arguments(s)>) like addNumbers(2, 4)

			PARAMETERS VS ARGUMENTS

			- Function parameters are the names listed in the function's definition
			- Function arguments are the real values passed into the function
*/

package main

import "fmt"

func main() {
	sumOfNumbers := addNumbers(2, 3)
	fmt.Println(sumOfNumbers)
	printGreeting("joe")
}

func addNumbers(a, b int) int {
	sum := a + b
	return sum
}

func printGreeting(str string) {
	fmt.Println("Hey there,", str)
}
