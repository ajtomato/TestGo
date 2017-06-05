package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"runtime"
	"strings"
	"sync"
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

	// Go has no comma operator and ++ and -- are statements not expressions.
	// Thus if you want to run multiple variables in a for you should use
	// parallel assignment (although that precludes ++ and --).
	a := []int{1, 2, 3}
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	fmt.Println(a)

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

	k := '?'
	switch k {
	case ' ', '?', '&', '=', '#', '+', '%':
		fmt.Println("Match")
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

func testRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
	for i := range pow {
		fmt.Printf("%d\n", pow[i])
	}
}

func testMap() {
	var m map[string]Vertex
	if m == nil {
		fmt.Println("A declared map is nil")
	}
	// Elements can NOT be inserted map without make
	m = make(map[string]Vertex)
	m["Hello"] = Vertex{1, 2}
	fmt.Println(m)

	var n = map[string]Vertex{
		"hello": Vertex{1, 2},
		"world": Vertex{3, 4},
	}
	fmt.Println(n)

	var l = map[string]Vertex{
		"hello": {1, 2},
		"world": {3, 4},
	}
	fmt.Println(l)
	delete(l, "hello")
	fmt.Println(l)
	elem, ok := l["hello"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("No hello elment", elem)
	}
	elem, ok = l["world"]
	if ok {
		fmt.Println(elem)
	} else {
		fmt.Println("No world elment")
	}
}

func (v *Vertex) abs() float64 {
	// It is common to write methods that gracefully handle being called with a
	// nil receiver.
	if v == nil {
		fmt.Println("v == nil")
		return 0.0
	}
	return math.Sqrt(float64(v.X*v.X + v.Y*v.Y))
}

func (v *Vertex) scale(s int) {
	v.X *= s
	v.Y *= s
}

func testMethod() {
	v := Vertex{3, 4}
	fmt.Println(v.abs())
	v.scale(2)
	fmt.Println(v)
}

type abser interface {
	abs() float64
}

func testInterface() {
	v := Vertex{3, 4}
	var a abser
	// a is NOT initialized, so it cannot point to any concrete type.
	fmt.Printf("%v\n", a)
	var tmp *Vertex
	// a is initialized, but the value is nil and there is a concrete type.
	a = tmp
	fmt.Printf("%v, %v, %T\n", a.abs(), a, a)
	// There is a concrete type for each interface value.
	a = &v
	fmt.Printf("%v, %v, %T\n", a.abs(), a, a)

	// The empty interface may hold values of any type.
	var i interface{}
	fmt.Printf("(%v, %T)\n", i, i)
	i = v
	fmt.Printf("(%v, %T)\n", i, i)

	// type assertion
	j, ok := i.(Vertex)
	if ok {
		fmt.Printf("(%v, %T)\n", j, j)
	} else {
		fmt.Println("i is NOT Vertex")
	}

	// type switch
	switch k := i.(type) {
	case Vertex:
		fmt.Printf("(%v, %T)\n", k, k)
	case int:
		fmt.Printf("(%v, %T)\n", k, k)
	default:
		fmt.Println("Unknown type")
	}
}

type myError struct {
	when time.Time
	what string
}

func (err myError) Error() string {
	return fmt.Sprintf("Error[%v]: %s\n", err.when, err.what)
}

func triggerError() error {
	return myError{time.Now(), "Huge Error"}
}

func testError() {
	err := triggerError()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("No error")
	}
}

type myReader struct {
	s io.Reader
}

func (r *myReader) Read(b []byte) (int, error) {
	c, err := r.s.Read(b)
	// Process b here
	return c, err
}

func testReader() {
	src := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := myReader{src}
	io.Copy(os.Stdout, &r)
}

func testGoroutine() {
	for i := 0; i < 5; i++ {
		go fmt.Println(i)
	}
	fmt.Println("testGoroutine")
	time.Sleep(100 * time.Millisecond)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func testChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func testBufferedChannel() {
	// Sends to a buffered channel block only when the buffer is full. Receives
	// block when the buffer is empty.
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Only the sender should close a channel
	close(c)
}

func closeChannel() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// The loop for i := range c receives values from the channel repeatedly
	// until it is closed.
	// v, ok := <-c; ok is false if there are no more values to receive and
	// the channel is closed.
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacciWithSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func testSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciWithSelect(c, quit)
}

func defaultSelection() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

type safeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *safeCounter) inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.v[key]++
}

func (c *safeCounter) get(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()

	return c.v[key]
}

func testMutex() {
	c := safeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.get("somekey"))
}

func namedResult(x, y int) (m, n int) {
	if x > y {
		m = x
		n = y
	} else {
		m = y
		n = x
	}
	return
}

func testNamedResult() {
	x, y := namedResult(2, 1)
	fmt.Printf("Big: %v, Small: %v\n", x, y)
	x, y = namedResult(7, 8)
	fmt.Printf("Big: %v, Small: %v\n", x, y)
}

type byteSize float64

const (
	_           = iota // ignore first value by assigning to blank identifier
	kB byteSize = 1 << (10 * iota)
	mB
	gB
	tB
	pB
	eB
	zB
	yB
)

func testConst() {
	fmt.Println(kB, mB, gB, tB, pB, eB, zB, yB)
}

func main() {
	// If an initializer is present, the type can be omitted. Please note that
	// c, python, java have different types.
	var f, python, java = 1.2, false, "no!"
	b = int(f)

	const w = "World"

	fmt.Printf(stringutil.Reverse("!oG ,olleH"))
	fmt.Printf("%v, %v, %v, %v, %v\n", c, python, java, e, d)

	testConst()
}
