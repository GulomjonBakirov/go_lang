package syntax

import "fmt"

func ChangeValue(x *int) {
	y := &x
	z := &y

	fmt.Printf("%v, %v, %v, %v\n", &x, x, &y, &z)

	*x = 100
}

func Pointer() {
	var num int = 10
	var ptr *int = &num // Pointer storing the address of num

	fmt.Println("Value: ", num)
	fmt.Println("Pointer address: ", ptr)
	fmt.Println("Pointer value: ", *ptr)
}
