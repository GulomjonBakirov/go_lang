package syntax

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func updateUserData(p *User) {
	p.Age += 1
}

// attached function to struct

func (p User) greet() {
	fmt.Println("Wussup ", p.Name)
}

func (p *User) increaseAge() {
	p.Age += 1
}

func Struct() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name) // Removes \n

	fmt.Print("Enter your age: ")
	var age int
	fmt.Scanln(&age)

	user := User{name, age}

	fmt.Printf("User Loaded: %+v\n", user)

	updateUserData(&user)

	fmt.Println(user)
	user.greet()
	user.increaseAge()
	fmt.Println(user.Age)

	jsonData, _ := json.Marshal(user)

	fmt.Println(string(jsonData))
}
