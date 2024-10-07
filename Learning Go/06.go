package main

func funcBuild() func(int, int) {
	return func(a, b int) {
		println(a + b)
	}
}

func main() {
	add := funcBuild()
	add(1, 2)
}
