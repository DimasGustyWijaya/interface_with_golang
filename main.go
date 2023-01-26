package main

import "fmt"

func main() {
	Dimas := Person{Name: "Dimas"}
	sayHe(Dimas)

	Cat := Animal{Name: "Cat"}
	sayHe(Cat)
}

// function Animal and his method
type Animal struct {
	Name string
}

func (a Animal) sayHi() string {
	return a.Name
}

// function Person and his method
type Person struct {
	Name string
}

func (p Person) sayHi() string {
	return "Hi" + p.Name
}

// interface and impl function nya
type HasName interface {
	sayHi() string
}

func sayHe(hasName HasName) {
	fmt.Println("hello", hasName.sayHi())
}

// nganu
