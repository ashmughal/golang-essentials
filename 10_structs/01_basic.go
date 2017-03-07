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

	fmt.Printf("%T, %v\n", person, person)

	person.age = 32
	person.firstName = "Ford"
	person.lastName = "Prefect"

	fmt.Printf("%T, %v\n", person, person)
}

