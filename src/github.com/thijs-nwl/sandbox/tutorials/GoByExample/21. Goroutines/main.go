package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 62; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine1")
	go f("goroutine2")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	fmt.Scanln()
	fmt.Println("done")
}
