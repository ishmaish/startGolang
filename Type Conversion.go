package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// converting String - > Integer
	// we use Atoi() function for that purpose
	// Ascii TO Integer -> ATOI
	// 01. ATOI
	strval := "100"
	intVar, err := strconv.Atoi(strval)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	// ParseInt interprets a string s in the given base(0,2 to 36) and bitsize (0,64)
	// and returns the value pf i
	// 02. ParseInt
	strval2 := "100"

	intVar2, err := strconv.ParseInt(strval2, 0, 8)
	fmt.Println(intVar2)

	// 03. using fmt.SScan
	strval3 := "100"
	intVal3 := 0
	_, err = fmt.Sscan(strval3, &intVal3)
	fmt.Println(intVal3, err, reflect.TypeOf(intVal3))

	// converting string - > float
	s := "3.1415926535"
	f, err := strconv.ParseFloat(s, 8)
	fmt.Println(f, err, reflect.TypeOf(f))

	// converting string - > Boolean
	s1 := "true"
	b1, _ := strconv.ParseBool(s1)
	fmt.Printf("%T, %v\n", b1, b1)

	// converting Boolean to String
	var b bool = true
	fmt.Println(reflect.TypeOf(b))

	var sam string = strconv.FormatBool(true)
	fmt.Println(reflect.TypeOf(sam))

	// Sprintf

	bed := true
	sed := fmt.Sprintf("%v", bed)
	fmt.Println(sed)
	fmt.Println(reflect.TypeOf(sed))

}
