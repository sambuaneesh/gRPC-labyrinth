package main

import "os"

func main() {

	createFile := func(fileName string) *os.File {
		file, _ := os.Create(fileName)
		defer file.Close()

		// write to file
		file.WriteString("Hello, World!")
		return file
	}

	openFile := func(fileName string) *os.File {
		file, _ := os.Open(fileName)
		return file
	}

	createFile("gg.txt")

	file := openFile("gg.txt")

	// read and print file
	fileContent := make([]byte, 100)
	file.Read(fileContent)
	println(string(fileContent))
}
