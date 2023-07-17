package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(name HasName) {
	fmt.Println("Hello", name.GetName())
}

type Test struct {
	Name string
}

func (person Test) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var husada Test
	husada.Name = "Husada"
	SayHello(husada)

	cat := Animal{
		Name: "Push",
	}
	SayHello(cat)
}
