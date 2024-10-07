package main

import "fmt"

func main() {

	iCanTakeAnything := func(i interface{}) {
		switch hisType := i.(type) {
		case int:
			fmt.Println("See I Took an int:", hisType)
		case string:
			fmt.Println("See I Took a string:", hisType)
		default:
			fmt.Println("See I Took something else:", hisType)
		}
	}

	iCanTakeAnything(1 + 2)
}
