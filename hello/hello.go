package main

import (
	"fmt"

	"ajtomato.gmail.com/test/stringutil"
)

// A var declaration can include initializers, one per variable.
var a, b int = 1, 2

var (
	c int
	d bool
	e string
)

// Constants cannot be declared using the := syntax.
const pi = 3.14

func main() {
	// If an initializer is present, the type can be omitted. Please note that
	// c, python, java have different types.
	var f, python, java = 1.2, false, "no!"
	b = int(f)

	const w = "World"

	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("%v, %v, %v, %v, %v", c, python, java, e, d)
}
