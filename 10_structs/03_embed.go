package main

import (
	"fmt"
)

type Person struct {
	age int
	firstName string
	lastName string
}

type Employee struct {
	Person

	id int
	firstName string
}

func main() {
	var employee Employee

	fmt.Printf("%T, %v\n", employee, employee)

	employee.age = 32
	employee.firstName = "Ford"
	employee.lastName = "Prefect"

	fmt.Printf("%T, %v\n", employee, employee)

	employee.age = 32
	employee.firstName = "Arthur"
	employee.Person.firstName = "Ford"
	employee.Person.lastName = "Prefect"

	fmt.Printf("%T, %v\n", employee, employee)
}

