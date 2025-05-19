package syntax

import "fmt"

func Loop() {
	var numbers []int = []int{1, 2, 3, 4, 5}

	for _, num := range numbers {
		fmt.Println(num)
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i)
	}
}
