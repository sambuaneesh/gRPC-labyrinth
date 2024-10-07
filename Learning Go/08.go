package main

import "fmt"

func main() {
	first := func() {
		fmt.Println("First")
	}

	second := func() {
		fmt.Println("Second")
	}

	third := func() {
		fmt.Println("Third")
	}

	goto ayoo

ayoo:
	first()
	second()
	goto byebye

byebye:
	third()
	goto ayoo
}
