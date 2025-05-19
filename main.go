package main

import (
	"fmt"
	"go_lang/syntax"
)

func main() {
	syntax.Variables()
	syntax.Loop()
	sum := syntax.AddSum(1, 2)
	success, result := syntax.Substract(4, 2)

	if success {
		fmt.Println(result)
	}

	fmt.Printf("Hello World %d \n", sum)
}
