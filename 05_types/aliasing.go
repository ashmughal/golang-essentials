package main

import (
	"fmt"
)

type foobar int16

func main() {
	var a foobar = 10
	var b int16 = 20

	//c := squareInt(a)
	//c := squareFoobar(b)
	c := squareFoobar(a)

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", c, c)
}

func squareInt(x int16) int16 {
	return x*x
}

func squareFoobar(x foobar) foobar {
	return x*x
}
