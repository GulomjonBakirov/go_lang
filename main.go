package main

import (
	"fmt"
	"go_lang/syntax"
)

// func main() {
// 	syntax.Variables()
// 	syntax.Loop()
// 	sum := syntax.AddSum(1, 2)
// 	success, result := syntax.Substract(4, 2)

// 	if success {
// 		fmt.Println(result)
// 	}

// 	fmt.Printf("Hello World %d \n", sum)
// }

func changeValue(n int) {
	n = 10

	fmt.Println(n)
}

func main() {
	number := 10

	syntax.Pointer()
	syntax.ChangeValue(&number)
	fmt.Println(number)

	changeValue(number)

	fmt.Println(number)

}
