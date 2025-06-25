/*
 *  src/hello-go/hello.go
 */

package main

import "fmt"

// A function can return any number of results.
// The `swap`` function returns two strings.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Printf("hello, world\n")

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}

/* EOF */
