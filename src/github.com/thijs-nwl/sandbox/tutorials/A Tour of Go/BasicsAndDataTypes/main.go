package main

import ( //importing packages
	"fmt"
	"math/rand"
	"time"
)

func add(x, y int) int { //costum function with parameters. type of parameters is the same so decalre it once
	return x + y
}

func swap(x, y string) (string, string) { //A fucntion can return any number of results
	return y, x
}

func split(sum int) (x, y int) { //Go's return values may be named. If so, they are treated as variables defined at the top of the function.
	x = sum * 4 / 9
	y = sum - x
	return //A return statement without arguments returns the named return values. This is known as a "naked" return.
}

var (
	c, python, java bool         //The var block statement declares a list of variables; as in function argument lists, the type is last. this is package level
	q, j            = 1, "nooo!" //var with initializer
)

var ( //vars declared without an value are given their zero value
	i int    //returns 0
	b bool   //returns false
	s string //returns ""(the empty string)
)

var x int = 45             //type conversion
var f float64 = float64(x) // The expression T(v) converts value v to tye T

const Pi = 3.14 //constant are declared like variables, but with the const keyword. Numeric constants are high-precision values

func main() { //main executing function
	z := 1e8
	k := 4                                              //function level var. Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.
	fmt.Printf("hello, world\n")                        //console logs in js
	fmt.Println("My favorite number is", rand.Intn(13)) //using of rand package
	fmt.Println("The time rn is", time.Now())

	fmt.Println(add(32, 32)) //costum fucntion call
	//fmt.Println(swap("hello", "world"))
	a, d := swap("hello", "world") //assign the return values of swap() to the variables a and b and print them
	fmt.Println(a, d)
	fmt.Println(split(45))
	fmt.Println(k, c, python, java, q, j, i, b, s)
	fmt.Println(z)
}
