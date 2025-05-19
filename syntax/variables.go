package syntax

import (
	"fmt"
)

// Global variables
var (
	firstName string = "G'ulomjon"
	lastName  string = "Bakirov"
	age       int    = 24
)

// global constants
const PI float64 = 3.14

func Variables() {
	var isAvailable bool = true
	var name string = "Bakirov"
	var age int = 25
	var pi float64 = 3.1415
	var character rune = 'Ж'
	var data byte = 255

	fmt.Println(isAvailable, name, age, pi, string(character), data)
}
