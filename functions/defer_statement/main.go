/*
A defer statement delays the execution of a function until the surrouding functions returns
the deferred call's are evaluated immediately, but the function call
is not executed until the surrounding function returns.
*/

package main

import "fmt"

func printName(str string) {
	fmt.Println(str)
}
func printRollNo(rno int) {
	fmt.Println(rno)
}

func printAddress(adr string) {
	fmt.Println(adr)
}

func main() {
	printName("Ramu")
	defer printRollNo(07)
	printAddress("NewYork")
}
