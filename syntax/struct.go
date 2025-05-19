package syntax

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	name string
	age  int
}

func updateUserData(p *User) {
	p.age += 1
}

// attached function to struct

func (p User) greet() {
	fmt.Println("Wussup ", p.name)
}

func (p *User) increaseAge() {
	p.age += 1
}

func Struct() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanln(&age)

	user := User{name, age}

	fmt.Printf("User Loaded: %+v\n", user)

	updateUserData(&user)

	fmt.Println(user)
	user.greet()
	user.increaseAge()
	fmt.Println(user.age)
}
