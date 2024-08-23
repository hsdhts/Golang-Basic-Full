package main

import "fmt"

func sayGoodBy(hello string) string {
	return "Hello" + hello
}

func main() {
	getSayGoodBy := sayGoodBy
	fmt.Println(getSayGoodBy("selamat pagi"))
}
