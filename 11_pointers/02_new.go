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
	person1 := new(Person)
	person2 := &Person{41, "Arthur", "Dent"}

	person1.age = 20
	person1.firstName = "Ford"
	person1.lastName = "Prefect"

	fmt.Printf("%T, %v\n", person1, person1)
	fmt.Printf("%T, %v\n", person2, person2)

	increaseAge(person1)
	increaseAge(person2)

	fmt.Printf("%T, %v\n", person1, person1)
	fmt.Printf("%T, %v\n", person2, person2)
}

func increaseAge(person *Person) {
	person.age += 10
}
