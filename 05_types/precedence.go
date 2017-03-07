package main

import (
	"fmt"
)

func main() {
	a :=  1+2*3-4/5%3
	b := (1+(2*3))-((4/5)%3)

	fmt.Println(a)
	fmt.Println(b)
}

