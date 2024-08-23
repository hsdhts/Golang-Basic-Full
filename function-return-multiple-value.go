package main

import "fmt"

func getFullName() (string, string) {
	return "Apa", "Aja"
}

func main() {
	firtsName, _ := getFullName()
	fmt.Println(firtsName)
}
