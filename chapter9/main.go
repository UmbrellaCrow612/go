package main

import (
	"fmt"
	"math"
)

type Circle struct { // Define our data type
	x float64
	y float64
	z float64
}

func (circle *Circle) area() float64 { // Way of writing a function called method
	return math.Pi * circle.x * circle.x // This allows us to invoke it with the . operator on the type you give it
}

func main() {
	circleOne := Circle{x: 1, y: 2, z: 3}
	fmt.Println(circleOne.area())

	a := new(Android)
	a.Person.Talk()
}

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person Person
	Model  string
}

/*

Think of it like object with property and methods both similar to javascript and c# classes

person = {
	name: "Tom",
	Talk: () => {
		console.log("Hello")
	}
}

person.Talk()


class Person {
	public name {get, set}

	public void Talk (){
		console.Write("Hello")
	}

	public person {
		this.name = name
	}
}


*/
