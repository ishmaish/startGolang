package main

import "fmt"

func main() {
	//Arithmetic Operators

	var x, y = 35, 7

	fmt.Printf("x + y = %d\n", x+y)
	fmt.Printf("x - y = %d\n", x-y)
	fmt.Printf("x * y = %d\n", x*y)
	fmt.Printf("x / y = %d\n", x/y)
	fmt.Printf("x mod y = %d\n", x%y)

	x++
	fmt.Printf("x++ = %d\n", x)

	y--
	fmt.Printf("y-- = %d\n", y)

	//Assignment Operators
	var x1, y1 = 15, 25
	x1 = y1
	fmt.Println("= ", x1)

	x1 = 15
	x1 += y1
	fmt.Println("+=", x1)

	x1 = 50
	x1 -= y1
	fmt.Println("-=", x1)

	x1 = 2
	x1 *= y1
	fmt.Println("*=", x1)

	x1 = 100
	x1 /= y1
	fmt.Println("/=", x1)

	x1 = 40
	x1 %= y1
	fmt.Println("%=", x1)

	//Comparison Operators
	var x2, y2 = 15, 25

	fmt.Println(x2 == y2)
	fmt.Println(x2 != y2)
	fmt.Println(x2 < y2)
	fmt.Println(x2 <= y2)
	fmt.Println(x2 > y2)
	fmt.Println(x2 >= y2)

	//Logical Operators
	var x3, y3, z3 = 10, 20, 30

	fmt.Println(x3 < y3 && x3 > z3)
	fmt.Println(x3 < y3 || x3 > z3)
	fmt.Println(!(x3 == y3 && x3 > z3))

	//Bitwise Operators
	var x4 uint = 9  //0000 1001
	var y4 uint = 65 //0100 0001
	var z4 uint

	z4 = x4 & y4
	fmt.Println("x & y  =", z4)

	z4 = x4 | y4
	fmt.Println("x | y  =", z4)

	z4 = x4 ^ y4
	fmt.Println("x ^ y  =", z4)

	z4 = x4 << 1
	fmt.Println("x << 1 =", z4)

	z4 = x4 >> 1
	fmt.Println("x >> 1 =", z4)
}
