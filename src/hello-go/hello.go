/*
 *  src/hello-go/hello.go
 */

package main

import (
	"fmt"
	"math/cmplx"
)

// Multiple results
//
// A function can return any number of results.
// The `swap` function returns two strings.
func swap(x, y string) (string, string) {
	return y, x
}

// Named return values
//
// Go's return values may be named. If so, they are treated as variables defined at the top of the function.
// These names should be used to document the meaning of the return values.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions, as with the example shown here.
// They can harm readability in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// Variables
//
// The `var` statement declares a list of variables; as in function argument lists, the type is last.
// A `var` statement can be at package or function level. We see both in this example.
var c, python, java bool

// Variables with initializers
//
// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var i, j int = 1, 2

// Short variable declarations
//
// Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.
// Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.
func shortVariableDeclarations() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}

// Basic types
//
// # Go's basic types are
//
// ```
// bool
//
// string
//
// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr
//
// byte // alias for uint8
//
// rune // alias for int32
//
//	represents a Unicode code point
//
// float32 float64
//
// complex64 complex128
// ```
//
// The example shows variables of several types, and also that variable declarations may be "factored" into blocks, as with import statements.
// The `int`, `uint`, and `uintptr` types are usually 32 bits wide on 32-bit systems and 64 bits wide on 64-bit systems.
// When you need an integer value you should use `int` unless you have a specific reason to use a sized or unsigned integer type.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func basicTypes() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

// Program entry point
func main() {
	fmt.Printf("hello, world\n")

	// Multiple results
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	// Named return values
	fmt.Println(split(17))

	// Variables
	var i int
	fmt.Println(i, c, python, java)

	// Variables with initializers
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	// Short variable declarations
	shortVariableDeclarations()

	// Basic types
	basicTypes()
}

/* EOF */
