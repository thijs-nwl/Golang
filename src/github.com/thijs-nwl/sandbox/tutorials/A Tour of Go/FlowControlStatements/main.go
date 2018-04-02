package main

import (
  "fmt"
  "math"
  "runtime"
  "time"
)

func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {     //Variables declared inside an if short statement are also available inside any of the else blocks.
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  return lim
}

func sqrt(x float64) string {
  if x < 0 {                            //Go if statements
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x))
}

func os(){
  fmt.Print("Go runs on ")
  switch os := runtime.GOOS; os {     //switch statements in go
    case "darwin":
      fmt.Println("OS X.")
    case "linux":
      fmt.Println("Linux.")
    case "windows":
      fmt.Println("windows.")
    default:
      fmt.Printf("%s.", os)
  }
}

func Saturday(){
  fmt.Println("When's Saturday?")   //Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
  today := time.Now().Weekday()
  switch time.Saturday {
  case today + 0:
    fmt.Println("Today.")
  case today + 1:
    fmt.Println("Tomorrow.")
  case today + 2:
    fmt.Println("In two days.")
  default:
    fmt.Println("Too far away.")
  }
}

func Hour(){                  //Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.
  t := time.Now()
  switch {
  case t.Hour() < 12:
    fmt.Println("Good morning!")
  case t.Hour() < 17:
    fmt.Println("Good afternoon")
  case t.Hour() < 24:
    fmt.Println("Good evening")
  default:
    fmt.Println("Good night")
  }
}

func Defers(){
  defer fmt.Println("world")    //The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

  fmt.Println("first")
  fmt.Println("hello")
  fmt.Println("then")


    fmt.Println("counting")       //Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
  	for i := 0; i < 10; i++ {
  		defer fmt.Println(i)
  	}
  	fmt.Println("done")

}

func main(){
  sum := 1                  //here you see C's while is spelled for in Go
  for sum < 1000 {
    sum += sum
  }
  fmt.Println(sum)

  fmt.Println(sqrt(2), sqrt(-4))

  fmt.Println(
    pow(3, 2, 10),
    pow(3, 3, 20),
  )

  os()
  Saturday()
  Hour()
  Defers()

  // sum := 1
  // for ; sum < 1000; {           //the init and post statement are optional
  //   sum += sum
  // }
  // fmt.Println(sum)

  // sum := 0
  // for i := 0; i < 10; i++ {           //Go has only one looping construct, the for loop.
  //   sum += i
  // }
  // fmt.Println(sum)

  // for {                      //if you omit the condition you get an infinite loop
  // }
}
