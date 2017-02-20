package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

func main() {
	var a complex64 = 1 + 2i
	var b complex64 = 2 - 2i
	var c complex64  // variable declared but not assigned gets the zero value
	var d complex128 = 3 + 5i
	e := 1 + 5i

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %T, %v\n", e, e)

	// addition
	fmt.Printf("%v + %v = %v\n", a, b, a + b)

	// subtraction
	fmt.Printf("%v - %v = %v\n", a, b, a - b)

	// multiplication
	fmt.Printf("%v * %v = %v\n", a, b, a * b)

	// division
	fmt.Printf("%v / %v = %v\n", a, b, a / b)

	// many operations on float64 are in the math package
	fmt.Printf("%v\n", cmplx.Sinh(math.Pi / 4.0))
}
