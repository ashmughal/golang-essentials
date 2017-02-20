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
	const a = true
	const b = false

	fmt.Print("\n")
	fmt.Println("boolConstants")
	fmt.Println("=============")

	fmt.Printf("%T, %v\n", a && b, a &&b)
}

func numericConstants() {
	const a = 100
	const b int = 100
	const c = 10.0
	var d float32 = a

	fmt.Print("\n")
	fmt.Println("numericConstants")
	fmt.Println("================")

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", c, c)
	fmt.Printf("%T, %v\n", d, d)
}

func runeConstants() {
	const a = 'a'
	const b rune = 'b'

	fmt.Print("\n")
	fmt.Println("runeConstants")
	fmt.Println("=============")

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
}

func stringConstants() {
	const a = "abc"
	const b string = "def"

	fmt.Print("\n")
	fmt.Println("stringConstants")
	fmt.Println("===============")

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
}

func complexNumberConstants() {
	const a = 1 + 2i
	const b complex64 = 2 + 3i
	const c complex64 = a

	fmt.Print("\n")
	fmt.Println("complexNumberConstants")
	fmt.Println("======================")

	fmt.Printf("%T, %v\n", a, a)
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", c, c)
}
