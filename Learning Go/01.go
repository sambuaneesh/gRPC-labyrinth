package main

import "fmt"

func main() {
	sayHello := func(name string) string {
		return "Hello " + name + "!!"
	}
	fmt.Println(sayHello("Go"))
}
