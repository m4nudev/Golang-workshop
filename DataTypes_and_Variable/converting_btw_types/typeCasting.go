package main

import (
	"fmt"
	"strconv"
)

func main() {
	//--------------------integer to float-----------------
	var i int = 90
	var f float64 = float64(i)
	fmt.Printf("%.2f \n", f)

	//--------------------integer to float-----------------
	var x float64 = 90.78
	var y int = int(x)
	fmt.Printf("%v \n", y)

	//--------------------strconv package-----------------
	/*
		Itoa()
				- converts integer to string
				- returns one value - string formed with the given integer.
	*/
	var integer int = 45
	var s string = strconv.Itoa(integer)
	fmt.Printf("%q\n", s)

	/*
		Atoi()
				- converts string to integer
				- returns two value - the corresponding integer, error (if any).
	*/
	var str string = "45"
	num, err := strconv.Atoi(str)
	fmt.Printf("%v, %T \n", num, num)
	fmt.Printf("%v, %T \n", err, err)

}
