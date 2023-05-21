package main

import "fmt"

// Go: Control Structures

func main() {

	// For control

	for i := 0; i < 10; i++ { // for statement allows us to repeat a list of statements
		fmt.Println(i)
	}

	// If control

	if 3%2 == 0 { // If statements check if the value is true then executes the code inside of it
		// even
	} else { // runs if the previous statement if false
		// odd
	}

	// Switch control

	switch 1 { // If the value matches a case then rub that
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	case 5:
		fmt.Println("Five")
	default:
		fmt.Println("Unknown Number")
	}

}
