package main

import (
	"fmt"
	"sync"
	"time"
)

//Goroutines
//A goroutine is a lightweight thread managed by the Go runtime
//starts a new goroutine running: go f(x,y,z)
//The evaluation of f, x, y and z happens in the current goroutine
//and the execution of f happens in the new goroutine
//Goroutines run in the same address space, so access to shared
//memory must be synchronized. The sync package provides useful
//primitives, although you won't need them much in go as there are
//other primitives
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func firstRoutine() {
	go say("world")
	say("hello")
}

//Channels
//Channels are a typed conduit through wich you can send and receive
//values with the channel operator, <-
// ch <- v			send v to channel ch
// v := <- ch		receive from ch, and assign value to v
//data flows in the direction of the arrow
//Like maps and slices, channels must be created before use: ch := make(chan int)
//By default, sends and receives block until the other side is ready.
//This allows goroutines to synchronize without explicit locks or condition variables
//The example code sums the numbers in a slice, distributing the work between
//two goroutines. Once both goroutines have completed their computation,
//it calculates the final result
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //send sum to c
}

//Buffered Channels
//Channels can be buffered. Provide the buffer length as the second
//argument to make to initialize a buffered channel: ch := make(chan int, 100)
//Sends to a buffered channel block only when the buffer is full.
//Receives block when the buffer is empty
func buffChan() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

//Range and Close
//A sender can close a channel to indicate that no more values will be send.
//Receivers can test whether a channel has been closed by assigning a second
//parameter to the receive expression: v, ok := <-ch
//ok is false if there are no more values to reveive and the channel is closed
//The loop for i:= range c receives values from the channel repeatedly until it is closed
//Note: Only the sender should close a channel, never the receiver.
//Sending on a closed channel will cause panic.

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

//Select
//The select statement lets a goroutine wait on multiple communication operations
//A selct blocks until one of its cases can run, then it executes that case.
//It chooses one at random if multiple are ready

func fibonacci2(c, quit chan int) {
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

//Default Selection
//the default case in a select is run if no other case is ready.
func TimeBomb() {
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

//sync.Mutex
//What if we just want to make sure onl one goroutine can access a variable at a time to
//avoid conflicts?
//This concept is called mutual exclusion, and the conventional name for the data structure
//that provides it is mutex
// mutex has two methods: lock and unlock

//safeCounter is safe to use concurrently
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

//Inc increments the counter for the given key
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	//Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

//Value returns the current value of the counter for the given key
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	//Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	//Goroutines
	firstRoutine()

	//Channels
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c //  recieve from c
	fmt.Println(x, y, x+y)

	//Range and Close
	ch := make(chan int, 100)
	go fibonacci(cap(ch), ch)
	for i := range ch {
		fmt.Println(i)
	}

	//Select
	ch2 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch2)
		}
		quit <- 0
	}()
	fibonacci2(ch2, quit)

	//default selection
	TimeBomb()

	//sync.Mutex
	sc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sc.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(sc.Value("somekey"))

}
