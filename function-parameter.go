package main

import "fmt"

func sayHelloto(firstName string, lastName string, umur int) {
	fmt.Println("Hello", firstName, lastName, umur)
}

func main() {
	sayHelloto("Husada", "Hutasoit", 21)
}
