package main

import (
	"fmt"
)

type Data struct {
	Name, Address string
	Age           int
}

func (customer Data) sayHi(name string) {
	fmt.Println("Hello", name, "My name is", customer.Name)
}

func (a Data) sayHuu() {
	fmt.Println("Huuu from", a.Name)
}

func main() {

	var husada Data
	husada.Name = " Husada"
	husada.Address = "Palembang"
	husada.Age = 21

	husada.sayHi("Joko")
	husada.sayHuu()
	//sayHi(husada, "Joko")

	//var husada Data
	//
	//husada.Name = "Husada"
	//husada.Address = "Palembang"
	//husada.Age = 21
	//
	//fmt.Println(husada)
	//
	//joko := Data{
	//	Name:    "Joko",
	//	Address: "Cirebon",
	//	Age:     20,
	//}
	//
	//fmt.Println(joko)
	//
	//budi := Data{"Budi", "Medan", 19}
	//
	//fmt.Println(budi)

}
