package main

import (
	"fmt"
	"reflect"
)

func main() {

	//---------------using %T----------------
	var grades int = 42
	var message string = "hello world"
	var isCheck bool = true
	var amount float32 = 4545.99

	fmt.Printf("variable grades = %v is of type %T \n", grades, grades)
	fmt.Printf("variable message = %v is of type %T \n", message, message)
	fmt.Printf("variable isCheck = %v is of type %T \n", isCheck, isCheck)
	fmt.Printf("variable amount = %v is of type %T \n", amount, amount)

	//---------------using reflect.TypeOf()----------------
	fmt.Printf("Type: %v \n", reflect.TypeOf(1000))
	fmt.Printf("Type: %v \n", reflect.TypeOf("KL Rahul"))
	fmt.Printf("Type: %v \n", reflect.TypeOf(78.0))
	fmt.Printf("Type: %v \n", reflect.TypeOf(true))

	var marks int = 67
	var greeting string = "Good morning"
	fmt.Printf("variable marks = %v is of type %v \n", marks, reflect.TypeOf(marks))
	fmt.Printf("variable marks = %v is of type %v \n", greeting, reflect.TypeOf(greeting))

}
