package main

import (
	"fmt"
)

// Constant block declaration
const (
	c_1 = 123
	c_2 = "four five six"
	c_3 = 7.89
)

// Enum definition with iota
// Iota resets in a new const block
const (
	v_1 = iota
	v_2 = iota
	v_3 // v_3 will implicitly get the next value of iota
	v_4 = iota
)

func do_types() {
	fmt.Println("Simple Types and Declarations")

	// Int, declaration separate from initialization
	var i int
	i = 42
	fmt.Println("i =", i)
	// Left shift - 2 << 42
	fmt.Println("2 << i =", 2<<i)
	fmt.Println()

	// Float, declaration and intialization
	// This is rounded at output to 7 decimal places, by float32?
	var f float32 = 3.1415926535897932384626433
	fmt.Println("f =", f)
	fmt.Println()

	// String, type inference
	s := "dreamt of"
	fmt.Println("s =", s)
	fmt.Println()

	// Pointers
	var s2 *string = new(string)
	*s2 = "correlate its contents"
	fmt.Println("s2 addr   =", s2)
	fmt.Println("s2 deref  =", *s2)
	fmt.Println()

	// Uninitialized pointer (<nil>)
	var ps *string
	fmt.Println("ps addr   =", ps)
	fmt.Println()

	// More pointer dereferencing
	s3 := "human mind"
	fmt.Println("s3 =", s3)
	ptr := &s3
	fmt.Println("ptr addr       =", ptr)
	fmt.Println("ptr deref      =", *ptr)
	s3 = "inability"
	fmt.Println("now, ptr addr  =", ptr)
	fmt.Println("now, ptr deref =", *ptr)
	fmt.Println()

	// Constants
	// Pi here is rounded to further precision (to 15 decimal places), maybe due
	// to type inferencing with const declarations?
	// Constants have to be evaluated at compile time, so no assignment of
	// return values of functions to constants.
	const pi = 3.1415926535897932384626433
	fmt.Println("pi =", pi)
	// Error - "cannot assign to pi"
	//pi = 6.28
	fmt.Println()

	// Implicitly-typed constants
	const c = 123
	// Here interpreted as int
	fmt.Println("c + 456  =", c+456)
	// Here interpreted as float
	fmt.Println("c + 78.9 =", c+78.9)
	fmt.Println()

	// To explicitly declare as an int:
	const c2 int = 123
	fmt.Println("c2 + 456  =", c2+456)
	// Now use as a float requires a cast:
	fmt.Println("c2 + 78.9 =", float32(c2)+78.9)
	fmt.Println()

	// The const values declared in the const block are available
	fmt.Println("c_2 =", c_2)
	fmt.Println("c_3 =", c_3)
	fmt.Println()

	// Enums - iota will increment by one each time it's used in a single
	// declaration
	fmt.Println("v_1 =", v_1)
	fmt.Println("v_2 =", v_2)
	fmt.Println("v_3 =", v_3)
	fmt.Println("v_4 =", v_4)
	fmt.Println()

	// Constant expressions and iota
	const (
		v_10 = iota - 1
		v_11 = iota * 20
		v_12 = 2 << iota
	)
	fmt.Println("v_10 =", v_10)
	fmt.Println("v_11 =", v_11)
	fmt.Println("v_12 =", v_12)
	fmt.Println()

	// Structs
	type struct_def struct {
		a int
		b float32
		c string
	}
	a_struct := new(struct_def)
	a_struct.a = 222
	a_struct.b = 6.62607
	a_struct.c = "alas, poor"
	fmt.Println("a_struct =", a_struct)
	fmt.Println()

	// Arrays - fixed size
	var arr1 [3]int
	arr1[0] = 10
	arr1[1] = 20
	arr1[2] = 30
	fmt.Println("arr1[1] =", arr1[1])
	//fmt.Println("arr1[3] =", arr1[3]) // Error - Index out of bounds
	fmt.Println("arr1 =", arr1)
	fmt.Println()

	// Arrays
	arr2 := [4]string{"one", "two", "three", "four"}
	fmt.Println("arr2 =", arr2)
	fmt.Println()

	// Slices
	slc1 := arr2[:]
	fmt.Println("slc1 =", slc1)
	// The slice is pointing to the array, so updates change both
	arr2[3] = "ten"
	slc1[0] = "eleven"
	fmt.Println("arr2 =", arr2)
	fmt.Println("slc1 =", slc1)
	fmt.Println()

	// Slices
	slc2 := []int{9, 8, 7, 6}
	fmt.Println("slc2 =", slc2)
	fmt.Println()
}

func main() {
	do_types()
}
