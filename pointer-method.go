package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr " + man.Name
}

func main() {
	husada := Man{"Husada"}
	husada.Married()

	fmt.Println(husada.Name)
}
