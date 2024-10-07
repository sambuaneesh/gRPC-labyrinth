package main

import "fmt"

// closures

func main() {
	counterBuilder := func() func() int {
		i := 0
		return func() int {
			i += 1
			return i
		}
	}

	ticketCounterA := counterBuilder()
	ticketCounterB := counterBuilder()

	fmt.Println(ticketCounterA())
	fmt.Println(ticketCounterA())
	fmt.Println(ticketCounterA())

	fmt.Println(ticketCounterB())
	fmt.Println(ticketCounterB())
}
