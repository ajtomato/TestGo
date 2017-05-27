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

func testDefer() {
	fmt.Println("counting")

	// Deferred function calls are pushed onto a stack.
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

// Vertex is for test.
type Vertex struct {
	X int
	Y int
}

func testStruct() {
	v := Vertex{1, 2}
	p := &v
	fmt.Println(v.X, p.Y)

	p = &Vertex{Y: 3}
	fmt.Println(*p)
}

func testSlice() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	// The capacity of a slice is the number of elements in the underlying
	// array, counting from the first element in the slice.
	s := primes[1:4]
	fmt.Println(s, len(s), cap(s))

	// Please note that s[1:4] is larger than the length of s.
	s = s[1:4]
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	s = s[1:]
	fmt.Println(s)
	s = s[:]
	fmt.Println(s)

	t := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(t)

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	w := make([]int, 4)
	y := make([]int, 3, 5)
	fmt.Println(w, len(w), cap(w))
	fmt.Println(y, len(y), cap(y))

	w = w[0:3]
	fmt.Printf("%p\n", &w[0])
	// Change the element in the underlying array
	w = append(w, 3)
	fmt.Printf("%p\n", &w[0])
	// Make a new array
	w = append(w, 3, 3, 3)
	fmt.Printf("%p\n", &w[0])
}

func main() {
	// If an initializer is present, the type can be omitted. Please note that
	// c, python, java have different types.
	var f, python, java = 1.2, false, "no!"
	b = int(f)

	const w = "World"

	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("%v, %v, %v, %v, %v\n", c, python, java, e, d)

	testSlice()
}
