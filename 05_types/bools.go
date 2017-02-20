package main

import "fmt"

func main() {
	var a bool = true  // variable declared and assigned the value true
	var b bool         // variable declared automatically gets the zero value type for boolean

	c := true          // RECOMMENDED: Declare and assign boolean value true
	d := false         // RECOMMENDED: Declare and assign boolean value true

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)

	fmt.Printf("> %v\n",  2 > 1)
	fmt.Printf("< %v\n",  2 < 1)
	fmt.Printf(">= %v\n",  2 >= 1)
	fmt.Printf("<= %v\n",  2 <= 1)

	fmt.Printf("a && b: %v\n", a && b)
	fmt.Printf("a || b: %v\n", a || b)
	fmt.Printf("a || b || d: %v\n", a || b || d)
}

