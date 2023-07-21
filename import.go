package main

import (
	"fmt"
	"golang-basic/helper"
)

func main() {
	helper.SayHello("name")
	//helper.sayGoodBye(18) error karena tidak bisa mengimpor menggunakan huruf kecil diawal
	fmt.Println(helper.SayHello)
}
