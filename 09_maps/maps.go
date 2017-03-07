package main

import (
	"fmt"
)

func main() {
	mapLiteral()
	makeMap()
	deleteElement()
	checkKeyExists()
}

func mapLiteral() {
	fmt.Println()
	fmt.Println("mapLiteral")
	fmt.Println("=========")

	a := map[string]string{"name": "Arthur Dent", "sex": "Male", "occupation": "BBC Radio Employee"}

	fmt.Printf("%T, %v, %v\n", a, a, len(a))
}

func makeMap() {
	fmt.Println()
	fmt.Println("makeMap")
	fmt.Println("=======")

	a := make(map[string]string, 10)
    var b map[string]string

	fmt.Printf("%T, %v, %v\n", a, a, len(a))
	fmt.Printf("%T, %v, %v\n", b, b, len(b))

	a["name"] = "Arthur Dent"
	//b["name"] = "Ford Prefect"
	fmt.Printf("%T, %v, %v\n", a, a, len(a))
	fmt.Printf("%T, %v, %v\n", b, b, len(b))
}

func deleteElement() {
	fmt.Println()
	fmt.Println("deleteElement")
	fmt.Println("=============")

	a := map[string]string{"name": "Arthur Dent", "sex": "Male", "occupation": "BBC Radio Employee"}
	fmt.Printf("%T, %v, %v\n", a, a, len(a))

	delete(a, "sex")
	fmt.Printf("%T, %v, %v\n", a, a, len(a))
}

func checkKeyExists() {
	fmt.Println()
	fmt.Println("checkKeyExists")
	fmt.Println("==============")

	a := map[string]string{"name": "Arthur Dent", "sex": "Male", "occupation": "BBC Radio Employee"}
	fmt.Printf("%T, %v, %v\n", a, a, len(a))

	v1, ok1 := a["sex"]
	v2, ok2 := a["dateOfBirth"]
	fmt.Printf("%T, %v, %v\n", v1, v1, ok1)
	fmt.Printf("%T, %v, %v\n", v2, v2, ok2)
}
