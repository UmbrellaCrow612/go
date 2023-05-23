package main

import "fmt"

// Pointers: will modify the original x variable in another function

// Pointers reference a location in memory where a value is stored rather than the value itself

// Whenever you need to modify a piece of data within a function you can with pointers

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1
}

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}
