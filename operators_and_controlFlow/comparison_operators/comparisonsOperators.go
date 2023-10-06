package main

import "fmt"

/*
compare two operands and yield a Booelan value
allows values of the same data type for comparisons
*/

func main() {
	// Equal to (==): returns true when the values are equal.

	var city string = "Kolkata"
	var city_2 string = "calcutta"
	fmt.Println(city == city_2)

	// Not Equal to (!=): returns true when the values are not equal.

	var city_3 string = "Kolkata"
	var city_4 string = "calcutta"
	fmt.Println(city_3 != city_4)

	// Less than ( < ): returns true when left operand is lesser than the right operand.

	var a, b int = 5, 10
	fmt.Println(a < b)

	// Less than equal to ( <= ): returns true when left operand is lesser or equal to the right operand.

	var x, y int = 10, 10
	fmt.Println(x <= y)

	// Greater than ( > ): returns true when left operand is greater than the right operand.

	var x1, y1 int = 11, 10
	fmt.Println(x1 > y1)

	// Greater than equal to ( >= ): returns true when left operand is greater or equal to the right operand.

	var x2, y2 int = 20, 20
	fmt.Println(x2 <= y2)

}
