package main

import (
	"fmt"
	"reflect"
)

func main() {
	// This program demonstrates the use of a variadic function in Golang. A variadic function is a function that takes a variable number of arguments of a specific type.

	//To declare a variadic function, the type of the final parameter is preceded by an ellipsis, "...", which shows that the function may be called with any number of arguments of this type.
	variadicExample("red", "blue", "green", "yellow")

	// normal func call

	fmt.Println(calculation("Rectangle", 20, 30))
	fmt.Println(calculation("Square", 20))

	// variadicexample2
	variadicExample2(1, "red", true, 10.5, []string{"foo", "bar", "baz"},
		map[string]int{"apple": 23, "tomato": 13})
}
func variadicExample(s ...string) {
	fmt.Println(s[0])
	fmt.Println(s[3])

	//but trying to access an index beyond the length of the slice would result in a runtime error.

}

// normal function parameter
func calculation(str string, y ...int) int {

	area := 1

	for _, val := range y {
		if str == "Rectangle" {
			area *= val
		} else if str == "Square" {
			area = val * val
		}
	}
	return area
}

// multiple type of arguments passed in variadic function

func variadicExample2(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v, "--", reflect.ValueOf(v).Kind())
	}
}
