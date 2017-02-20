package main

import (
	"fmt"
)

func main() {
	singleDimension()
	multiDimensional()
	genericArray()

	// arrays are value objects
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	arrayValueParam(a)
	fmt.Printf("a[0]=%v\n", a[0])
}

func singleDimension() {
	var a [5]int                          // array of size with each value initialized to zero
	var b [5]int = [5]int{1, 2, 3, 4, 5}  // array initialized from an array literal
	c := [5]int{1, 2, 3, 4, 5}

	fmt.Println()
	fmt.Println("singleDimension")
	fmt.Println("===============")

	fmt.Printf("%T, %v, len=%v, cap=%v\n", a, a, len(a), cap(a))
	fmt.Printf("%T, %v, len=%v, cap=%v\n", b, b, len(b), cap(b))
	fmt.Printf("%T, %v, len=%v, cap=%v\n", c, c, len(c), cap(c))

	fmt.Println("===============")

	for i := 0; i < len(a); i++ {
		fmt.Println(i, b[i])
	}

	fmt.Println("===============")

	for i, v := range b {
		fmt.Println(i, v)
	}
}

func multiDimensional() {
	var a [5][5]string = [5][5]string{[5]string{"a", "b", "c", "d", "e"}}

	fmt.Println()
	fmt.Println("multiDimensional")
	fmt.Println("================")

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%v\n", a[0][0])
}

func genericArray() {
	var a [5]interface{} = [5]interface{}{1, true, 1 + 2i, 10.0, "abc"}

	fmt.Println()
	fmt.Println("genericArray")
	fmt.Println("================")

	for i, v := range a {
		fmt.Printf("%v, %T, %v\n", i, v, v)
	}
}

func arrayValueParam(a [5]int) {
	fmt.Printf("a[0]=%v\n", a[0])
	a[0] = 25
	fmt.Printf("a[0]=%v\n", a[0])
}
