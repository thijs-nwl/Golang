package main

import (
	"fmt"
	"io"
	"math"
	"strings"
	"time"
)

type Vertex struct {
	X, Y float64
}

type MyFloat float64

//Methods
//A method is a function with a special receiver argument.
//The receiver appears in its own argument list between the func keyword and the method name.
//In this example, the Abs method has a receiver of type Vertex named v.
//Remember: a method is just a function with a receiver argument.
//You can only declare a method with a receiver whose type is
//defined in the same package as the method. You cannot
//declare a method with a receiver whose type is defined in
//another package (which includes the built-in types such as int).
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

//Pointer receivers
//You can declare methods with pointer receivers.
//This means the receiver type has the literal syntax *T for some
//type T. (Also, T cannot itself be a pointer such as *int.)
//With a value receiver, the Scale method operates on a copy of the original Vertex value.
//(This is the same behavior as for any other function argument.) The Scale method must have a
//pointer receiver to change the Vertex value declared in the main function.
//Choosing a value or pointer receiver
//The first is so that the method can modify the value that its receiver points to.
//The second is to avoid copying the value on each method call. This can be more efficient
//if the receiver is a large struct, for example.
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Interfaces
//A type implements an interface by implementing its methods.
//Under the covers, interface values can be thought of as a tuple of a value and a concrete type: (value, type)
//An interface value holds a value of a specific underlying concrete type.
//Calling a method on an interface value executes the method of the same name on its underlying type.
//If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.
type I interface {
	M()
}

type T struct {
	h, w string
}

type F float64

//This method means type T implements the interface I,
//but we don't need to explicitly declare that it does so
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.h, t.w)
}

func (f F) M() {
	fmt.Println(f)
}

//Empty interface
//The interface type that specifies zero methods is known as the empty interface:
//interface{}
//An empty interface may hold values of any type. (Every type implements at least zero methods.)
func emptyI() {
	var i interface{}
	fmt.Println(i)

	i = 42
	fmt.Println(i)

	i = "hello"
	fmt.Println(i)
}

//Type assertions
//A type assertion provides access to an interface value's underlying concrete value.
//t := i.(T)
//this statement asserts that te interface value i holds te concrete type T and assigns
//the underlying T value to the variable t.
//to test whether an interface vlaue holds a specific type, a type assertion can return two values:
//the underlying value and a boolean value that reports whether the assertion succeeded
//t, ok := i.(T)
//if i holds a T, then t will be the underlying value and ok will be true
//if not, ok will be false and t will be the zero value of type T, and no panic occurs
func typeAssertions() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64) //No panic
	fmt.Println(f, ok)

	// f = i.(float64) PANIC
	// fmt.Println(f)
}

//Type switches
//A type switch is a construct that permits several type assertions in series
//A type switch is like a regular switch statement, but the cases in a type switch specify
//types (not values), and those values are compared against the type of the value held
//by the given interface value.
func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice int %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("What the fack is a %T!\n", v)
	}
}

//Stringer
//A Stringer is a type that can describe itself as a string
//The fmt package (and many others) look for this interface to print values
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

//Errors
//Go programs express error state with error values.
//the error type is a built-in interface similar to fmt.Stringer
type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

//Readers
//The io package specifies the io.Reader interface, which represents the
//read end of a stream of data
//The Go standard library contains many implementations of these interfaces,
//including files, network connections, compressors, chipher, and others

func main() {
	//methods
	v := Vertex{3, 4}
	fmt.Println(v.Abs())

	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(10)
	fmt.Println(v.Abs())

	//interfaces
	var i I

	var t *T
	i = t
	i.M()

	i = &T{"hello", "world"}
	i.M()

	i = F(math.Pi)
	i.M()

	//empty interface
	emptyI()

	//type assertions
	typeAssertions()

	//type switches
	typeSwitch(21)
	typeSwitch("hello")
	typeSwitch(true)

	//Stringer
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	//Error
	if err := run(); err != nil {
		fmt.Println(err)
	}

	//Reader
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
}
