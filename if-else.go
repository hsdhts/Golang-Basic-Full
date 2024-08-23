package main

import "fmt"

func main() {

	name := "Husada"

	if name == "Husada" {
		fmt.Println("Hello Husada")
	} else if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Nama salah")
	}

	//if wtih short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama tidak valid")
	} else {
		fmt.Println("Nama valid")
	}
}
