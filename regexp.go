package main

import (
	"fmt"
	regexp2 "regexp"
)

func main() {

	var regexp = regexp2.MustCompile("e([a-z])o")

	fmt.Println(regexp.MatchString("eko"))
	fmt.Println(regexp.MatchString("budi"))
	fmt.Println(regexp.MatchString("hasan"))

	search := regexp.FindAllString("eko, edo, eto, eyo, eki", -1)

	fmt.Println(search)
}
