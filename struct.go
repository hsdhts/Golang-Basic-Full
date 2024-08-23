package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int64
}

func (customer Customer) sayAja(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name, "umur saya", customer.age, "tahun")
}

func main() {
	var customer Customer
	customer.Name = "Husada"
	customer.Address = "Palembang"
	customer.age = 21

	fmt.Println(customer)

	//masukkan kedalam variable
	budi := Customer{"Budi", "US", 22}
	fmt.Println(budi)

	//Struct Method
	budi.sayAja("Cindy")
}
