package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct {
	Name string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

func (p Person) GetName() string {
	return p.Name
}

func main() {
	person := Person{"Husada"}
	SayHello(person)
}
