package main

import "fmt"

func getComplete() (firstName, middleName, lastName string) {
	firstName = "Husada"
	middleName = "MK"
	lastName = "HTS"

	return firstName, middleName, lastName
}

func main() {
	a, b, c := getComplete()
	fmt.Println(a, b, c)
}
