package main

import (
	"fmt"
	"os"
)

func main() {
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// 	return
	// }

	file := new(os.File)

	if _, err := os.Stat("test.txt"); os.IsNotExist(err) {
		file, _ := os.Create("test.txt")
		fmt.Fprintln(file, "Bonjour le monde")
	} else {
		file, _ = os.Open("test.txt")
	}
	// defer file.Close() // defer postpones the execution of the function until the surrounding function returns

	// open the file for reading
	file, _ = os.Open("test.txt")
	// read and store the last line of the file in variable str
	var str string
	// go to the last line of the file
	fmt.Fscanln(file, &str)

	// convert the string to a slice of bytes
	bs := []byte(str)
	// print the slice of bytes
	fmt.Println(bs)
	fmt.Println(str)

	// loop through bs and increment each byte by 1
	for i, v := range bs {
		bs[i] = v + 1
	}
	fmt.Println(bs)

	file.Close()

	file, _ = os.Open("test.txt")
	// convert the slice of bytes back to a string
	str = string(bs)
	fmt.Println(str)
	file, _ = os.Create("test.txt")
	fmt.Fprintln(file, str)
	file.Close()

}
