package helper

import "fmt"

var Application = "Backend Golang"

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodBye(age int) {
	fmt.Println("Umur saya", age)
}
