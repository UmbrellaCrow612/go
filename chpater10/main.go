package main

import "fmt"

// Concurrency: Making progress on more than one task simultaneously is known as concurrency.

func f(n int) {

	for i := 0; i < 10; i++ {

		fmt.Println(n, ":", i)

	}
}

func main() {

	// Goroutine: is a function that is capable of running concurrently with other functions

	go f(0)

	// Channels: Channels provide a way for two goroutines to communicate with one another and synchronize their execution

	var c chan string = make(chan string)

	fmt.Println(c)

}
