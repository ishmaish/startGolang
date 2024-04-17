package main

import "fmt"

//Go has a special statement called defer that schedules a function call to be run after the function completes.
//A defer statement is often used
//with paired operations like open and
//close, connect and disconnect, or
//lock and unlock to ensure that resources are released in all cases, no matter how complex the control flow. The right place
//for a defer statement that releases a resource is immediately after the resource
//has been successfully acquired.

func main() {
	// the program will print second

	defer second()
	first()
}

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}

// Readwrite func with defer

/*func ReadWrite() bool {
	file.Open("file")
	defer file.Close()   //file.Close() is added to defer list
	// Do your thing
	if failureX {
		return false   // Close() is now done automatically
	}

	if failureY {
		return false   // And here too
	}

	return true   // And here
}*/

/*It keeps our Close call near our Open call so it's easier to understand.
If our function had multiple return statements (perhaps one in an if and one in an else), Close will happen before both of them.
Deferred Functions are run even if a runtime panic occurs.
Deferred functions are executed in LIFO order, so the above code prints: 4 3 2 1 0.
*/
