package main

import (
	"fmt"
	"runtime"
	"time"

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

func loop() {
	sum := 0
	// for i := 0, j := true is NOT allowed
	for i, j := 0, true; i < 10 && j; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// loops forever
	// for {
	// }
}

func testSwitch() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}

	// Switch cases evaluate cases from top to bottom, stopping when a case
	// succeeds.
	i, j := 0, 3
	switch j {
	case i + 3:
		fmt.Println("switch case 1")
		i = 3
	case i:
		fmt.Println("switch case 2")
	}

	// long if-then-else
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func main() {
	// If an initializer is present, the type can be omitted. Please note that
	// c, python, java have different types.
	var f, python, java = 1.2, false, "no!"
	b = int(f)

	const w = "World"

	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("%v, %v, %v, %v, %v\n", c, python, java, e, d)

	testSwitch()
}
