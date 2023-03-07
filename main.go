package main

// factored import statement
import (
	"fmt"
	"io"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
	"time"
)

/*
import "fmt"
import "math"
*/
var c, python, java bool
var i, j int = 1, 2

const Pii = 3.14
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// Struct is a collection of fields
type Vertex struct {
	X int
	Y int
}

func main() {
	var i int = 40
	fmt.Println("My favorite number is", rand.Intn(100), i)
	fmt.Println(add(42, 13))
	fmt.Println(add1(46, 13, 56))
	fmt.Println(swapReturnTuple("12", "14"))
	fmt.Println(split(100))

	var c, python, java = true, false, "no!"
	str := "denememee"
	fmt.Println(i, j, c, python, java, str)

	log()

	zeroValues()

	typeConversion()

	typeInference()

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	forLoop()
	switchCase()
	deferExample()
	pointers()
	structOps()

	mapExample()

	functionValues()

	myMethodExample()

}

func add(x int, y int) int {
	return x + y
}

func add1(x, y, z int) int {
	return x + y + z
}

func swapReturnTuple(x, y string) (string, string) {
	return y, x
}

/*
Go's return values may be named. If so, they are treated as variables defined at the top of the function.
These names should be used to document the meaning of the return values.
A return statement without arguments returns the named return values. This is known as a "naked" return.
Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
*/
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func log() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

func typeConversion() {
	var x, y int = 3, 4
	fmt.Println(x, y)
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func typeInference() {
	var i int
	j := i // j is an int

	k := 42           // int
	f := 3.142        // float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf("j is of type %T\n", j)
	fmt.Printf("k is of type %T\n", k)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("g is of type %T\n", g)
}

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func forLoop() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
	/*
		infinite loop
		for {
			}
	*/
}

func ifCondition(x int) {
	if x < 0 {
		fmt.Println(x)
	}
}

func pow(x, n, lim float64) float64 {
	//if with short statement
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return v + 1
	}
	return lim
}

func switchCase() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switchWithNocondition() {
	/*
		Switch without a condition is the same as switch true.
		This construct can be a clean way to write long if-then-else chains.
	*/
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

func deferExample() {
	/*
		The deferred call's arguments are evaluated immediately,
		but the function call is not executed until the surrounding function returns.
	*/
	i := 0
	defer fmt.Println(i)
	for ; i < 100; i++ {

	}
	fmt.Println(i)

	/*
		Deferred function calls are pushed onto a stack.
		When a function returns, its deferred calls are executed in last-in-first-out order.
	*/
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

func pointers() {
	i, j := 42, 2701

	p := &i            // point to i
	fmt.Println(*p, j) // read i through the pointer
	*p = 21            // set i through the pointer
	fmt.Println(i)     // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func structOps() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

	structPointer()
}

func structPointer() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

	// that notation is cumbersome, so the language permits us instead to write just p.X,
	// without the explicit dereference.
	(*p).Y = 12
	fmt.Println(v)

	structLiterals()
}

func structLiterals() {
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p  = &Vertex{1, 2} // has type *Vertex
	)

	fmt.Println(v1, p, v2, v3)
	v3 = Vertex{13, 13}
	fmt.Println(v1, p, v2, v3)
	arrays()
}

func arrays() {
	// An array's length is part of its type, so arrays cannot be resized.
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// Slices
	var s []int = primes[1:4]
	fmt.Println(s)

	s[2] = 99
	s[0] = 99
	fmt.Println(primes)

	sliceLiteral()
}

func sliceLiteral() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
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

	fmt.Println(s)
	sliceDefaults()
}

func sliceDefaults() {
	// When slicing, you may omit the high or low bounds to use their defaults instead.
	// The default is zero for the low bound and the length of the slice for the high bound.

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Printf("len=%d cap=%d %v\n", len(primes), cap(primes), primes)

	s := primes[1:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = primes[:2]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = primes[1:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	dynamicSizedArray()
}

func dynamicSizedArray() {
	println("------")
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)

	i := append(a, 999, 99)
	printSlice("i", i)

	rangeLoop()
}

func printSlice(s string, x []int) {

	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func rangeLoop() {
	println("-----")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

}

func mapExample() {
	println("-----")
	//var m map[string]Vertex
	m := make(map[string]Vertex)

	m["deneme"] = Vertex{1, 1}
	m["deneme1"] = Vertex{2, 2}

	fmt.Println(m)

	mapLiterals()
}

func mapLiterals() {
	println("-----")
	m := map[string]int{"a": 12, "b": 353}
	fmt.Println(m)

	mutatingMap()
}

func mutatingMap() {
	println("------")
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 2
	m["c"] = 3

	delete(m, "b")
	v, u := m["c"]
	fmt.Println(m)
	fmt.Println("value: ", v, "present", u)
}

func functionValues() {
	println("-------")
	myMultiplyFunction := func(x, y int) int {
		return x * y
	}

	myAddFunction := func(x, y int) int {
		return x + y
	}

	println(myMultiplyFunction(3, 4))
	println(myAddFunction(3, 4))

	println("------")
	println(funcWithFuncArgument(myMultiplyFunction, 12, 12))
	println(funcWithFuncArgument(myAddFunction, 12, 12))

	println("function closures")
	myFn := returnsFuncClosure()

	println(myFn(23, 34))
}

func funcWithFuncArgument(functionArg func(int, int) int, x, y int) int {
	res := functionArg(x, y)
	return res
}

func returnsFuncClosure() func(int, int) int {
	sum := 0

	return func(x int, y int) int {
		sum += x + y
		return sum
	}

}

// A method is a function with a special receiver argument.
func myMethodExample() {
	println("---------")
	vertex := Vertex{4, 5}
	fmt.Println(vertex.myMethod())
	fmt.Println(vertex)
	vertex.myMethodPointer(2)
	fmt.Println(vertex.myMethod())
	fmt.Println(vertex)

	var i = MyFloat(12.3)
	fmt.Println(i.mySecondMethod())

	// interface
	var k I = &T{"hello"}
	k.M()

	interfaceWithNull()
}

func (receiver Vertex) myMethod() int {
	receiver.X = receiver.X * 2
	receiver.Y = receiver.Y * 2
	return receiver.X + receiver.Y
}

func (receiver *Vertex) myMethodPointer(scale int) {
	receiver.X = receiver.X * scale
	receiver.Y = receiver.Y * scale
}

type MyFloat float64

func (receiver MyFloat) mySecondMethod() float64 {
	if receiver < 0 {
		return float64(-receiver)
	}
	return float64(receiver * 2)
}

type I interface {
	M()
}

type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// in case of nil value, method is invoked instead of throwing null pointer exception
func interfaceWithNull() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()

	typeAssertion()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func typeAssertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println("string- " + s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// if i carries float, copy value to f and set ok as true.
	f, ok := i.(float64)
	fmt.Println(f, ok)

	//f = i.(float64) // panic
	//fmt.Println(f)

	typeSwitch()
}

func typeSwitch() {
	do(21)
	do("hello")
	do(true)

	stringer()
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

// Stringers
// The fmt package (and many others) look for this interface to print values.
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("__ %v (%v years)", p.Name, p.Age)
}

func stringer() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	testError()
}

// Errors
// As with fmt.Stringer, the fmt package looks for the error interface when printing values.
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

// A nil error denotes success; a non-nil error denotes failure.
func testError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	readerExample()
}

// readers
func readerExample() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}

	typeParameterExample()
}

// Type parameters
// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func typeParameterExample() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

	runGoroutine()
}

// List represents a singly-linked list that holds
// values of any type.
// Generic type
type List[T any] struct {
	next *List[T]
	val  T
}

// A goroutine is a lightweight thread managed by the Go runtime.
// Goroutines run in the same address space, so access to shared memory must be synchronized.
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func runGoroutine() {
	go say("world")
	say("hello")

	channelExample()
}

// Channels
// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// By default, sends and receives block until the other side is ready.
// This allows goroutines to synchronize without explicit locks or condition variables.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// sum = <-c // kanaldaki veri bu şekilde alınır
	c <- sum // send sum to c
}

func channelExample() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	bufferedChannel()
}

// Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.
func bufferedChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	rangeAndCloseExample()
}

// A sender can close a channel to indicate that no more values will be sent.
// Channels aren't like files; you don't usually need to close them.
// Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
func rangeAndCloseExample() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Print(i, " ")
	}

	selectExample()
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case.
// It chooses one at random if multiple are ready.
func selectExample() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci1(c, quit)
}

func fibonacci1(c, quit chan int) {
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
