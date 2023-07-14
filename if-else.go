package main

import (
	"fmt"
)

func main() {

	name := "Husada"

	if name == "Husada" {
		fmt.Println("Hello Husada")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Masukkan nama yang benar!")
	}

	//if with short statement
	if length := len(name); length > 5 {
		fmt.Println("Karakter nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
