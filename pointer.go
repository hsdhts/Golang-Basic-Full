package main

import "fmt"

type Address struct {
	City, Name, Province string
}

func main() {
	address1 := Address{"Bandung", "Ordent", "Jabar"}
	address2 := &address1 // pointer

	address2.City = "Medan"
	fmt.Println(address1)
	fmt.Println(address2)
}
