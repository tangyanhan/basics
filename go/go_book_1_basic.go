package main;
// This is a basic practice for go language
/*
	By Ethan Tang
	Copyright Ethan Tang 2017-2017
*/


import "fmt"

func main(){
	fmt.Println("Hello World!")

	var a float64 = 1.1
	var b float64 = 0.9
	var c = a - b
	fmt.Println(a, "-", b, "=", c)

	// Integers. Go has interger of type int/uint8/uint32/uint64/int8/int32/int64, etc
	// Variables of different types cannot be computed directly,
	// types must be explicitly converted when necessary
	var ia int32 = 11
	ib := 22
	fmt.Println(ia, "+", ib, "=", ia + int32(ib))

	// Boolean types
	ba := true
	bb := false
	fmt.Println(ba, "&&", bb, "=", (ba&&bb))
	fmt.Println(ba, "||", bb, "=", (ba||bb))
	fmt.Println("!", ba, "=", !ba)

	// String operations
	var sa string = "Hello "
	sb := "Ethan"
	fmt.Println("len(", sa, ")=", len(sa))
	fmt.Println(sa, "+", sb, "=", sa+sb)

	// Declare multiple variables
	// A variable in Go must be initialized before being used.
	var (
		mv_a = 1;  // Go statement may end with ; or leave it alone
		mv_b = 2
		mv_c = 3
	)
	fmt.Println(mv_a, mv_b, mv_c)

	// Define a constant
	// := is used to declare variable so it should not be used on constants
	const cst_str = "I'm a constant string"
	const cst_ano_str string = "I'm another constant string"
	const cst_int = 11
	const cst_ano_int int = 12
	fmt.Println("Constants:", cst_str, cst_ano_str, cst_int, cst_ano_int)
}
