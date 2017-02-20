package main

import (
	"fmt"
	"math"
)

func main() {
	var a float32 = math.SmallestNonzeroFloat32
	var b float32 = math.MaxFloat32
	var c float32                    // variable declared but not assigned gets the zero value
	var d float64 = math.SmallestNonzeroFloat64
	var e float64 = math.MaxFloat64
	var f float64                    // variable declared but not assigned gets the zero value
	g := 20.2

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("d: %v\n", e)
	fmt.Printf("d: %v\n", f)
	fmt.Printf("d: %T, %v\n", g, g)

	// addition
	fmt.Printf("%v + %v = %v\n", 20.1, 10.2, 20.1 + 10.2)

	// subtraction
	fmt.Printf("%v - %v = %v\n", 20.1, 10.2, 20.1 - 10.2)

	// multiplication
	fmt.Printf("%v * %v = %v\n", 20.1, 10.2, 20.1 * 10.2)

	// division
	fmt.Printf("%v / %v = %v\n", 20.1, 10.2, 20.1 / 10.2)

	// many operations on float64 are in the math package
	fmt.Printf("%v, %v\n", math.Sin(math.Pi/4.0), 1/math.Sqrt(2))
}
