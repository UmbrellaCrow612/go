package main

import "fmt"

// Go: Variables and control flow statements

// A variable is a storage location, with a specific type and an associated name identifier

func main() {

	var x string // Declare a variable string

	x = "Hello World" // assign the variable a value "x is assigned the string Hello World"

	fmt.Println(x)

	x = "Hello World 2"

	fmt.Println(x) // variables can change their value throughout the lifetime of a program

	x = x + " special assignment statement" // Adds the current value to a new one and then assigns that to the variable

	fmt.Println(x)

	x += " Shorthand assignment" // Same as above

	fmt.Println(x == "") // Compares the value and returns a boolean

	// Declaring variables

	y := "" // infer the type based on the literal value you assign the variable
	var z = ""
	var zz string = ""
	const zzz = ""

	fmt.Println(y+z+zz, zzz)

	useThisWay := "" // Generally you should use this shorter form whenever possible.

	fmt.Println(useThisWay)

	// Scope: Variables inside a function cannot be accessed by those above them, functions inside of the original can.

	// Go Constants: Constants are basically variables whose values cannot be changed later.

	const stayTheSame string = "I can not be changed"

	// Defining multiple variables at once

	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println(a, b, c)

}
