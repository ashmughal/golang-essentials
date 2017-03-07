package main

import (
	"fmt"
)

func main() {
	boolConstants()
	numericConstants()
	runeConstants()
	stringConstants()
	stringConstants()
	complexNumberConstants()
}

func boolConstants() {
	fmt.Println()
	fmt.Println("boolConstants")
	fmt.Println("=============")

	const a = true
	const b = false

	fmt.Printf("%T, %v\n", a && b, a &&b)
}

func numericConstants() {
	fmt.Println()
	fmt.Println("numericConstants")
	fmt.Println("================")

	const a = 100
	const b int = 100
	const c = 10.0
	var d float32 = a

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("%T, %v\n", d, d)
}

func runeConstants() {
	fmt.Println()
	fmt.Println("runeConstants")
	fmt.Println("=============")

	const a = 'a'
	const b rune = 'b'

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
}

func stringConstants() {
	fmt.Println()
	fmt.Println("stringConstants")
	fmt.Println("===============")

	const a = "abc"
	const b string = "def"

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
}

func complexNumberConstants() {
	fmt.Println()
	fmt.Println("complexNumberConstants")
	fmt.Println("======================")

	const a = 1 + 2i
	const b complex64 = 2 + 3i
	const c complex64 = a

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", c, c)
}
