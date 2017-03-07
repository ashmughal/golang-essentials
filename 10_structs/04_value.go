package main

import (
	"fmt"
)

type Person struct {
	age int
	firstName string
	lastName string
}

func main() {
	var person Person

	person.age = 21
	person.firstName = "Ford"
	person.lastName = "Prefect"

	fmt.Printf("%T, %v\n", person, person)

	increaseAge1(person)

	fmt.Printf("%T, %v\n", person, person)

	increaseAge2(&person)

	fmt.Printf("%T, %v\n", person, person)
}

// person is passed by value
func increaseAge1(person Person) {
	person.age += 10

	fmt.Printf("increaseAge1: %T, %v\n", person, person)
}

// person is passed by reference
func increaseAge2(person *Person) {
	person.age += 10

	fmt.Printf("increaseAge2: %T, %v\n", person, person)
}
