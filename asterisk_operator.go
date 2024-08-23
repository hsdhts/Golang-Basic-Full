package main

import "fmt"

type Data struct {
	City, Name, Province string
}

func main() {
	addresss1 := Data{"Bandung", "Ordent", "Jabar"}
	addresss2 := &addresss1

	addresss2.City = "Jakarta"
	fmt.Println(addresss1)
	fmt.Println(addresss2)

	*addresss2 = Data{"Medan", "Backend", "Sumut"}
	fmt.Println(addresss2)

}
