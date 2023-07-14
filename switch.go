package main

import "fmt"

func main() {

	name := "Budi"

	switch name {
	case "Husada":
		fmt.Println("Hello Husada")
	case "Budi":
		fmt.Print("Hello Budi")
	default:
		fmt.Print("Masukkan nama yang benar")
	}

	//switch with short statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Karakter nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah sesuai")
	}

}
