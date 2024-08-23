package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func changeCountry(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{"Bandung", "Ordent", ""}
	changeCountry(&address)
	fmt.Println(address)
}
