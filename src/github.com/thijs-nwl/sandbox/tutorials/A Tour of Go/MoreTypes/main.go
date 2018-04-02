package main

import (
	"fmt"
	"math"
)

type Vertex struct { //A struct is a collection of fields.
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	t  = &Vertex{1, 2} // has type *Vertex
)

func pointers() {
	fmt.Println("pointers()")
	defer fmt.Println("\n")
	i, j := 42, 2701

	p := &i         //point to i
	fmt.Println(*p) //read i through the pointer
	*p = 21         //set i through the pointer
	fmt.Println(i)  //see the new value of i

	p = &j         //point to j
	*p = *p / 37   //divide j through the pointer
	fmt.Println(j) //see the new value of j

	v := Vertex{1, 2}
	v.X = 4 //Struct fields are accessed using a dot.
	fmt.Println(v.X)

	z := &v // pointer to the v Structs
	z.X = 1e9
	fmt.Println(v)

	fmt.Println(v1, v2, v3, t)
}

func arr() {
	fmt.Println("arr()")
	defer fmt.Println("\n")
	var a [2]string //The type [n]T is an array of n values of type T
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} //An array's length is part of its type, so arrays cannot be resized.
	fmt.Println(primes)
}

func slice() {
	fmt.Println("slice()")
	defer fmt.Println("\n")
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //This selects a half-open range which includes the first element, but excludes the last one.
	fmt.Println(s)
}

func Xslice() {
	fmt.Println("Xslice()")
	defer fmt.Println("\n")
	names := [4]string{ //A slice does not store any data, it just describes a section of an underlying array.
		"John", //Changing the elements of a slice modifies the corresponding elements of its underlying array.
		"Paul", //Other slices that share the same underlying array will see those changes.
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"

	fmt.Println(a, b)
	fmt.Println(names)
}

func sliceLiteral() {
	fmt.Println("sliceLiteral()")
	defer fmt.Println("\n")
	q := []int{2, 3, 5, 7, 11, 13} //A slice literal is like an array literal without the length.
	fmt.Println(q)                 //This is an array literal:
	//[3]bool{true, true, false}
	r := []bool{true, false, true, true, false, true} //And this creates the same array as above, then builds a slice that references it:
	fmt.Println(r)                                    //[]bool{true, true, false}

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
	fmt.Println(s[0])

	x := []int{2, 3, 5, 7, 11, 13}

	x = x[1:4]
	fmt.Println(x)

	x = x[:2]
	fmt.Println(x)

	x = x[1:]
	fmt.Println(x)

	x = x[:]
	fmt.Println(x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func sliceInfo() {
	fmt.Println("sliceInfo()")
	defer fmt.Println("\n")
	s := []int{2, 3, 5, 7, 11, 13} //The length of a slice is the number of elements it contains.
	printSlice(s)                  //The capacity of a slice is the number of elements in the underlying array, counting
	//from the first element in the slice.
	//The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
	//Slice the slice to give it zero length
	s = s[:0]
	printSlice(s)

	//Extend its length
	s = s[:4]
	printSlice(s)

	//Drop its first two values
	s = s[2:]
	printSlice(s)
}

func makeSlice() {
	fmt.Println("makeSlice()")
	defer fmt.Println("\n")
	a := make([]int, 5)
	printMakeSlice("a", a)

	b := make([]int, 0, 5)
	printMakeSlice("b", b)

	c := b[:2]
	printMakeSlice("c", c)

	d := c[2:5]
	printMakeSlice("d", d)
}

func printMakeSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func appendSlice() {
	fmt.Println("appendSlice()")
	defer fmt.Println("\n")
	var s []int
	printSlice(s)

	//append works on nil slices
	s = append(s, 0)
	printSlice(s)

	//the slice grows as needed
	s = append(s, 1)
	printSlice(s)

	//we can add more than one elem at a time
	s = append(s, 2, 3, 4)
	printSlice(s)

	//more about slices here:
	//https://blog.golang.org/go-slices-usage-and-internals
}

func rangeOver() {
	fmt.Println("rangeOver()") //When ranging over a slice, two values are returned for each iteration.
	defer fmt.Println("\n")    //The first is the index, and the second is a copy of the element at that index.
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

func mappings() { //A map maps keys to values.
	fmt.Println("mappings()") //The make function returns a map of the given type, initialized and ready for use.
	defer fmt.Println("\n")   //m = make(map[string]Vertex)

	type Vertex struct {
		Lat, Long float64
	}

	m := map[string]Vertex{ //Map literals are like struct literals, but the keys are required.
		"Bell Labs": {40.68433, -74.39967}, //If the top-level type is just a type name, you can omit it from the elements of the literal.
		"Google":    {37.42202, -122.08408},
	}

	fmt.Println(m)
}

func MutatingMaps() {
	fmt.Println("MutatingMaps()")
	defer fmt.Println("\n")
	m := make(map[string]int)

	m["Answer"] = 42 //instert or update an elem in map m: m[key] = elem
	fmt.Println("the value:", m["Answer"])

	m["Answer"] = 48 //retrieve an elem: elem = m[key]
	fmt.Println("the value:", m["Answer"])

	delete(m, "Answer") //delete an elem: delete(m, key)
	fmt.Println("the value:", m["Answer"])

	v, ok := m["Answer"] //test that a key is present with a two-value assignment: elem, ok = m[key]        ok = true || false
	fmt.Println("The value:", v, "Present?", ok)
}

func functionValues() {
	fmt.Println("functionValues()")
	defer fmt.Println("\n")
	compute := func(fn func(float64, float64) float64) float64 {
		return fn(3, 4)
	}
	//Functions are values too. They can be passed around just like other values.Function values may be used as function arguments and return values.
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(3, 4))
	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

//Function closures
//Go functions may be closures. A closure is a function value that references variables from outside its body.
//The function may access and assign to the referenced variables; in this sense the function is "bound"
//to the variables.

//For example, the adder function returns a closure. Each closure is bound to its own sum variable.

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func Exec() {
	fmt.Println("Exec()")
	defer fmt.Println("\n")
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}

func main() {
	pointers()
	arr()
	slice()
	Xslice()
	sliceLiteral()
	sliceInfo()
	makeSlice()
	appendSlice()
	rangeOver()
	mappings()
	MutatingMaps()
	functionValues()
	Exec()
}
