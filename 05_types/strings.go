package main

import (
	"fmt"
)

func main() {
	a := "你好世界"  // hello world in Chinese
	b := `a multiline string
	  which stretches onto the next line and contains embedded "quotations" marks
		`
	var c string  // variable declared but not assigned gets the zero value
	d := "the quick brown fox ..."

	fmt.Printf("a: %v, %v, %v\n", len(a), len([]rune(a)), a)
	fmt.Printf("b: %v, %v\n", len(b), b)
	fmt.Printf("c: %v, %v\n", len(c), c)
	fmt.Printf("d: %v, %v\n", len(d), d)
	fmt.Printf("d: %v, %v\n", len(a + b), a + b)

	fmt.Printf("d: %v, %v\n", a[0], a[1])

	runeOperations()
	byteOperations()
}

func runeOperations() {
	fmt.Println()
	fmt.Println("runeOperations")
	fmt.Println("================")

	a := "你好世界"  // hello world in Chinese
	b := '好'

	fmt.Printf("%T, %v, %v", b, b, string(b))

	for _, r := range a {
		fmt.Printf("%v, %v\n", r, string(r))
	}
}

func byteOperations() {
	fmt.Println()
	fmt.Println("byteOperations")
	fmt.Println("================")

	a := "你好世界"  // hello world in Chinese
	b := []byte{228, 189, 160, 229, 165, 189, 228, 184, 150, 231, 149, 140}

	fmt.Printf("%T, %v\n", a[0], a[0])
	fmt.Printf("=%v\n", []byte(a))
	fmt.Printf("%v, %v\n", a, string(b))
}
