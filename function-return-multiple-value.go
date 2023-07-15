package main

import "fmt"

func getFullName() (string, string, string) {
	return "Husada", "Karunia", "Hutasoit"
}

func main() {
	firstName, _, lastName := getFullName()
	fmt.Println(firstName)
	//fmt.Println(middleName) contoh: untuk ignore value yang tidak digunakan
	fmt.Println(lastName)
}
