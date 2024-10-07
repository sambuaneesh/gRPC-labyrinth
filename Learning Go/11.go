package main

import "fmt"

func main() {
	growSomeBalls := func(currBalls int, grownBalls chan int) {
		grownBalls <- currBalls + 1
	}

	yourBalls := make(chan int)

	go growSomeBalls(0, yourBalls)
	go growSomeBalls(69, yourBalls)
	go growSomeBalls(-1, yourBalls)
	go growSomeBalls(9999, yourBalls)

	fmt.Println(<-yourBalls, <-yourBalls, <-yourBalls, <-yourBalls)

}
