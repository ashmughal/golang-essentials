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
	person := Person{20, "Ford", "Prefect"}

	fmt.Printf("%T, %v\n", person, person)

	increaseAge(&person)

	fmt.Printf("%T, %v\n", person, person)
}

func increaseAge(person *Person) {
	person.age += 10
}
