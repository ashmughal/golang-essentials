package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceLiteral()
	makeSlice()
	convertSlice()

	// slice operatins
	sliceAppend()
	copySlice()
	subSlice()
	deleteElement()
	sortSlice()
}

func sliceLiteral() {
	fmt.Println()
	fmt.Println("sliceLiteral")
	fmt.Println("=========")

	a := []int{1, 2, 3, 4, 5}

	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))
}

func makeSlice() {
	fmt.Println()
	fmt.Println("makeSlice")
	fmt.Println("=========")

	a := make([]int, 10, 20)

	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))
}

func convertSlice() {
	fmt.Println()
	fmt.Println("convertSlice")
	fmt.Println("============")

	a := "你好世界" // hello world in chinese
	b := [5]int{1, 2, 3, 4, 5}

	a1 := []byte(a)
	a2 := []rune(a)
	b1 := b[:]

	fmt.Printf("%T, %v, %v, %v\n", a1, a1, len(a1), cap(a1))
	fmt.Printf("%T, %v, %v, %v\n", a2, a2, len(a2), cap(a2))
	fmt.Printf("%T, %v, %v, %v\n", b1, b1, len(b1), cap(b1))
}

func sliceAppend() {
	fmt.Println()
	fmt.Println("sliceAppend")
	fmt.Println("===========")

	var a []int   // declaration without assignment sets to zero value of type which is Nil for slices

	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))

	a = append(a, 1)
	a = append(a, 2, 3, 4, 5)

	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))
}

func copySlice() {
	fmt.Println()
	fmt.Println("copySlice")
	fmt.Println("=========")

	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))

	b := make([]int, 3)
	fmt.Printf("%T, %v, %v, %v\n", b, b, len(b), cap(b))

	copy(b, a)  // note: destination is the first parameter !
	fmt.Printf("%T, %v, %v, %v\n", b, b, len(b), cap(b))
}

func subSlice() {
	fmt.Println()
	fmt.Println("subSlice")
	fmt.Println("========")

	a := []int{1, 2, 3, 4, 5}
	b := a[:2]
	c := a[2:]
	d := a[:]

	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))
	fmt.Printf("%T, %v, %v, %v\n", b, b, len(b), cap(b))
	fmt.Printf("%T, %v, %v, %v\n", c, c, len(c), cap(c))
	fmt.Printf("%T, %v, %v, %v\n", d, d, len(d), cap(d))
}

func deleteElement() {
	fmt.Println()
	fmt.Println("deleteElement")
	fmt.Println("=============")

	a := []int{1, 2, 3, 4, 5}
	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))

	a = append(a[:2], a[3:]...)
	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))
}

func sortSlice() {
	fmt.Println()
	fmt.Println("sortSlice")
	fmt.Println("=========")

	a := []int{5, 4, 3, 2, 1}
	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))

	sort.Ints(a)
	fmt.Printf("%T, %v, %v, %v\n", a, a, len(a), cap(a))
}
