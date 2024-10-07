package main

import "fmt"

// func main() {
// 	iCanTakeAll := func(params ...int) int {
// 		sum := 0
// 		for _, param := range params {
// 			sum += param
// 		}
// 		return sum
// 	}

// 	println(iCanTakeAll(1, 2, 3, 4, 5))
// }

func main() {
	iCanTakeAllShit := func(params ...any) []any {
		disList := []any{}
		for _, param := range params {
			// fmt.Printf("%T\n", param)
			disList = append(disList, param)
		}
		return disList
	}

	iCanReduceShit := func(disList []any) int {
		reducedShit := 0
		for _, item := range disList {
			if val, isInt := item.(int); isInt {
				reducedShit += val
			} else {
				if str, isString := item.(string); isString {
					stringByteShit := []byte(str)
					for _, shit := range stringByteShit {
						reducedShit += int(shit)
					}
				}
			}
		}

		return reducedShit
	}

	heresMyShit := iCanTakeAllShit(1, "hello", "shit", 0)
	seeMyShit := iCanReduceShit(heresMyShit)
	fmt.Println(seeMyShit)
}
