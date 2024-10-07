package main

import "fmt"

func main() {
	bs := []byte("a slice") // []byte("a slice") is a conversion
	// demonstrate before and after conversion by printing the type
	// of the variable bs
	fmt.Printf("Before conversion: %T\n", bs)
	fmt.Printf("After conversion: %T\n", string(bs))
	fmt.Println(bs[2])
	fmt.Printf("%c\n", bs[2])
	fmt.Printf("%q\n", bs[2])
	fmt.Printf("%T", bs[2])

	s := []int{1, 2}
	_ = s // if i am not using the var but i want to keep

	cookin := []int{1, 2}
	print(cookin, "\n")
	fmt.Println(cookin)
	cookin = append(cookin, 3)
	fmt.Println(cookin)
	cookin = append(cookin, []int{4, 5, 6}...)
	fmt.Println(cookin)
	print(len(cookin), "\n")
}
