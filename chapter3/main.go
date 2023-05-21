package main

import "fmt"

// Go types - store a piece of data depending on what it is

// Numbers: In go you split them into whole integers and floating point numbers i.e decimals.

// Go's integer types are: uint8, uint16, uint32, uint64, int8, int16, int32 and int64.

// The numbers after uint are the bits each type has, unit means unassigned integer, int means assigned integer

// Generally if you are working with integers you should just use the int type.

// Go has two floating point types: float32 and float64

// go has additional types for representing complex numbers (numbers with imaginary parts): complex64 and complex128

// Generally we should stick with float64 when working with floating point numbers.

func main() {

	fmt.Println("1 + 1 =", 1+1)

	fmt.Println("1 + 1 =", 1.0+1.0)

	stringExample()

	booleanExample()

}

/*


Operators in go


+	addition
-	subtraction
*	multiplication
/	division
%	remainder

*/

// Strings: Go strings are made up of individual bytes, usually one for each character

// String literals can be created using double quotes "Hello World" or back ticks `Hello World`.

// The difference between these is that double quoted strings cannot contain newlines and they allow special escape sequences.

func stringExample() {
	fmt.Println(len("Hello World")) // A space is also considered a character
	fmt.Println("Hello World"[1])   // Strings are “indexed” starting at 0 not 1, think of them as arrays of individual characters
	fmt.Println("Hello " + "World") // Concatenation uses the same symbol as addition
}

func booleanExample() {

	/*

		Three logical operators are used with boolean values

			&&	and
			||	or
			!	not

	*/

	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

}

// at https://www.golang-book.com/books/intro/3
