package main

import "fmt"

func main() {
	/*
		syntax : var <Array name> [<size of the array>] <data type>
		var grades [5] int
		var fruits [3] string

		var grades [3]int = [3]int{10,20,30}
		shorthand method : [3]int{10,20,30}
		grades := [...]int{10, 20, 30}
	*/

	var fruits [2]string = [2]string{"apples", "oranges"}
	fmt.Println(fruits)
	fmt.Println(len(fruits))

	marks := [3]int{10, 20, 30}
	fmt.Println(marks)
	fmt.Println(len(marks))

	names := [...]string{"rachel", "phoebe", "Monica"}
	fmt.Println(names)
	fmt.Println(len(names))

	/*
		len() : The length of the array refers to the number of elements stored in the array.
	*/

	// how to access the element of the array

	var fruits1 [5]string = [5]string{"apples", "oranges", "grapes", "mango", "papaya"}
	fmt.Println(fruits1[4])

}
