package main

import (
	"fmt"
	"math"
)

func main() {
	signedIntegers()
	unsignedIntegers()
	integerOperators()
	bitwiseOperators()
}

func signedIntegers() {
	var a int8 = math.MaxInt8     // int8 ranges from -128 to 127
	var b int16 = math.MaxInt16   // int16 ranges from -32768 to 32767
	var c int = math.MaxInt64     // int ranges from -9223372036854775808 to 9223372036854775807 (on a 64 bit processor)
	var d int32 = math.MaxInt32   // int32 ranges from -2147483648 to 2147483647
	var e int64 = math.MaxInt64   // int64 ranges from -9223372036854775808 to 9223372036854775807 (on a 64 bit processor)

	fmt.Print("\n")
	fmt.Println("signedIntegers")
	fmt.Println("==============")

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

	fmt.Print("\n")
	fmt.Println("Widening conversions")
	fmt.Println("====================")
	fmt.Printf("int8 -> int16: %v, %v\n", a, int16(a))
	fmt.Printf("int16 -> int32: %v, %v\n", b, int32(b))
	fmt.Printf("int32 -> int64: %v, %v\n", d, int64(d))

	fmt.Print("\n")
	fmt.Println("Narrowing conversions")
	fmt.Println("=====================")
	fmt.Printf("int16 -> int8: %v, %v\n", b, int8(b))
	fmt.Printf("int32 -> int16: %v, %v\n", d, int16(d))
	fmt.Printf("int64 -> int32: %v, %v\n", e, int32(e))
}

func unsignedIntegers() {
	var a uint8 = math.MaxUint8     // unsigned int8 ranges from 0 to 255
	var b uint16 = math.MaxUint16   // unsigned int16 ranges from 0 to 65535
	var c uint = math.MaxUint64     // unsigned int ranges from 0 to 18446744073709551615 (on a 64 bit processor)
	var d uint32 = math.MaxUint32   // unsigned int32 ranges from 0 to 4294967295
	var e uint64 = math.MaxUint64   // unsigned int64 ranges from 0 to 18446744073709551615

	fmt.Print("\n")
	fmt.Println("unsignedIntegers")
	fmt.Println("================")

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

	fmt.Print("\n")
	fmt.Println("Widening conversions")
	fmt.Println("====================")
	fmt.Printf("uint8 -> uint16: %v, %v\n", a, uint16(a))
	fmt.Printf("uint16 -> uint32: %v, %v\n", b, uint32(b))
	fmt.Printf("uint32 -> uint64: %v, %v\n", d, uint64(d))

	fmt.Print("\n")
	fmt.Println("Narrowing conversions")
	fmt.Println("=====================")
	fmt.Printf("uint16 -> uint8: %v, %v\n", b, uint8(b))
	fmt.Printf("uint32 -> uint16: %v, %v\n", d, uint16(d))
	fmt.Printf("uint64 -> uint32: %v, %v\n", e, uint32(e))
}

func integerOperators() {
	a := 101
	b := 10
	c := 9223372036854775807

	fmt.Print("\n")
	fmt.Println("integerOperators")
	fmt.Println("================")

	// addition
	fmt.Printf("%v + %v = %v\n", a, b, a + b)

	// subtraction
	fmt.Printf("%v - %v = %v\n", a, b, a - b)

	// multiplication
	fmt.Printf("%v * %v = %v\n", a, b, a * b)

	// division
	fmt.Printf("%v / %v = %v\n", a, b, a / b)

	// modulus
	fmt.Printf("%v %% %v = %v\n", a, b, a % b)

	// overflow
	fmt.Printf("%v * %v = %v\n", c, 10, c * 10)
}

func bitwiseOperators() {
	fmt.Print("\n")
	fmt.Println("bitwiseOperators")
	fmt.Println("================")

	// bit wise and
	// 10     = 1010
	// 3      = 0011
	// 10 & 3 = 1011 = 11
	fmt.Printf("%b | %b: %v\n", 10, 3, 10 | 3)

	// bit wise and
	// 10     = 1010
	// 3      = 0011
	// 10 & 3 = 0010 = 2
	fmt.Printf("%b & %b: %v\n", 10, 3, 10 & 3)

	// bit wise xor
	// 10     = 1010
	// 3      = 0011
	// 10 ^ 3 = 1001 = 9
	fmt.Printf("%b ^ %b: %v\n", 10, 3, 10 ^ 3)

	// bit clear
	// 10      = 1010
	// 3       = 0011
	// 10 &^ 3 = 1001 = 9
	fmt.Printf("%b &^ %b: %v\n", 10, 3, 10 &^ 3)

	// bit shift left
	// 2      = 0010
	// 2 << 2 =  1000 = 8
	fmt.Printf("%v << %v: %v\n", 2, 2, 2 << 2)

	// bit shift right
	// 2      = 1090
	// 8 >> 2 = 0010 = 2
	fmt.Printf("%v >> %v: %v\n", 8, 1, 8 >> 1)
}
