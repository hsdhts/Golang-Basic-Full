package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Husada",
		"address": "Palembang",
	}
	person["title"] = "Software Engineer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)
	book["title"] = "Belajar Go-lang"
	book["author"] = "Husada"
	book["wrong"] = "Ups salah!"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "wrong")
	fmt.Println(book)
	fmt.Println(len(book))

}
