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
	arthur := Person{23, "Arthur", "Dent"}
	ford := Person{lastName: "Ford", firstName: "Prefect", age: 43}

	fmt.Printf("%T, %v\n", arthur, arthur)
	fmt.Printf("%T, %v\n", ford, ford)
}

