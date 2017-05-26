package main

import (
	"fmt"

	"ajtomato.gmail.com/test/stringutil"
)

// A var declaration can include initializers, one per variable.
var a, b int = 1, 2

func main() {
	// If an initializer is present, the type can be omitted. Please note that
	// c, python, java have different types.
	var c, python, java = true, false, "no!"

	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("%v, %v, %v", c, python, java)
}
