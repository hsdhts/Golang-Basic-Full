package main

import "fmt"

func main() {

	name := "Husada"

	switch name {
	case "Husada":
		fmt.Println("Hello Husada")
	case "Eko":
		fmt.Println("Hello Eko")
	case "Budi":
		fmt.Println("Hello Budi")
	default:
		fmt.Println("Nama tidak ditemukan")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama lumayan panjang")
	case false:
		fmt.Println("Nama valid")
	}
}
