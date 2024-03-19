package main

import (
	"fmt"
	"reflect"
)

// VARIABLES IN GO//
func main() {
	// TO BE INITIALIZED USING "VAR" KEYWORD
	// var name type = expression
	// variable declaration with initialization, hence Golang is strong satic typing it needs to specify the type of variable
	var i int = 10
	var s string = "Germany"

	fmt.Println(i, s)

	// 01: variable without initialization

	// METHOD:01

	var i1 int
	var s2 string

	i = 10
	s = "Germany"

	fmt.Println(i1, s2)

	// at the point of i1 is of type int is just declared but not assigned to meaningful values but later
	// they are assigned to values

	// 02: Variable with type inference

	// omitting the declaring part and the variable understands the type of variable based on the value type

	// INFERRED VARIABLE

	var i3 = 10
	var s3 = "Germany"

	fmt.Println(i3, s3)
	fmt.Println(reflect.TypeOf(i3))

	// 03: Multiple variable declaration
	var fname, lname string = "John", "Doe"
	m, n, o := 1, 2, 3

	fmt.Println(fname + lname)
	fmt.Println(m + n + o)

	// 04: SHORT VARIABLE DECLARATION

	// using this we don't have to use the var keyword or declare variable type

	name := "ishma faraj"
	fmt.Println(reflect.TypeOf(name))
}
