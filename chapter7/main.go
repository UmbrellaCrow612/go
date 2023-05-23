package main

import "fmt"

// Functions

func main() {

	// A function is an independent section of code that maps zero or more input parameters to zero or more output parameters.

	// Parameters func(here) are defined as name then type

	xs := []float64{98, 93, 77, 82, 83}

	fmt.Println(total(xs))

	defer second() // schedules a function call to be run after the function completes
}

func total(listOfTestScores []float64 /* Special prop ...name allows as many to be passed*/) float64 /*Return type if defined*/ {

	var total float64 = 0

	for _, v := range listOfTestScores {
		total += v
	}

	return total // Can return multiple values

	// You can also have closures: functions inside of functions

	// You can also have Recursion: running the function again

}

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func panicAndRecover() {
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
