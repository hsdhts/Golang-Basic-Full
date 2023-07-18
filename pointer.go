package main

import "fmt"

type Address struct {
	Name, City, Country string
}

func changeCountryToIndonesia(address *Address) {
	address.Country = "Jepang"
}

func main() {
	var Address1 Address = Address{"Husada", "Sumsel", "Indonesia"}
	var Address2 *Address = &Address1
	var Address3 *Address = &Address1

	Address2.City = "Bandung"
	Address2.Name = "Budi"

	Address2 = &Address{"Dimas", "London", "England"}

	fmt.Println(Address1)
	fmt.Println(Address2)
	fmt.Println(Address3)

	var Address4 *Address = new(Address)
	Address4.City = "Jakarta"
	fmt.Println(Address4)

	var alamat = Address{
		Name:    "Panji",
		City:    "Subang",
		Country: "",
	}

	changeCountryToIndonesia(&alamat)
	fmt.Println(alamat)
}
