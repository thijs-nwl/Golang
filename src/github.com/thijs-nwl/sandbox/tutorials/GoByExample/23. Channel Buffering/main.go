package main

import (
	"fmt"
)

func main() {
	messagges := make(chan string, 2)

	messagges <- "bufferd"
	messagges <- "channel"

	fmt.Println(<-messagges)
	fmt.Println(<-messagges)

}
