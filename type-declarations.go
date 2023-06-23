package main

import "fmt"

func main() {
	type NoKTP string
	type married bool

	var noKTPHusada NoKTP = "12345678"
	var marriedStatus married = true

	fmt.Println(noKTPHusada)
	fmt.Println(marriedStatus)
}
