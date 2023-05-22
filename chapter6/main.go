package main

import (
	"fmt"
)

// Go: Arrays, Slices and Maps

func main() {

	// Arrays: An array is a numbered sequence of elements of a single type with a fixed length, there length cannot be changed

	// Arrays start at index 0

	var x [5]int // This is how you declare an empty array of length 5 in go: You give it a name, size and type

	// Like many other languages, you have primitive functions you can apply to arrays.

	// For example: funding an index, length , adding, deleting and more

	// Type conversion, is invoked like a function

	// Special way od doing a for loop

	/*

			for _, value := range x {
		        total += value
		    }

	*/

	// the _ is used when you declare a variable that is not used and want a placeholder, the go compiler ignores this.

	y := [5]float64{98, 93, 77, 82, 83} // Short hand for making an array

	fmt.Println(x, y)

	// Slices: A slice is a segment of an array, Unlike arrays this length is allowed to change.

	var slice []float64 // The only difference between this and an array is the missing length between the brackets

	wayToMakeSlice := make([]float64, 5, 10) // you should use the built-in make function to create slices

	// Slices can never be longer than the array, they can be smaller

	// First parameter is the array and its type

	// The second is the length of the slice

	// The third is the capacity

	fmt.Println(slice, wayToMakeSlice)

	// Maps: A map is an unordered collection of key-value pairs

	// The map type is represented by the keyword map, followed by the key type in brackets and finally the value type

	mapExample := make(map[int]int) // If you were to read this out loud you would say “x is a map of strings to ints.”

	mapExample[1] = 10

	fmt.Println(mapExample, mapExample[1])

	elements := map[string]map[string]string{
		"H": {
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": {
			"name":  "Helium",
			"state": "gas",
		},
		"Li": {
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": {
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": {
			"name":  "Boron",
			"state": "solid",
		},
		"C": {
			"name":  "Carbon",
			"state": "solid",
		},
		"N": {
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": {
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": {
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": {
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println(el["name"], el["state"])
	}

	// We now have a map of strings to maps of strings to strings. The outer map is used as a lookup table based on the element's symbol, while the inner maps are used to store general information about the elements.

}
