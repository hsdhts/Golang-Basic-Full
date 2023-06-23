package main

import "fmt"

func main() {

	var result = 10 + 10
	fmt.Println(result)

	var a = 10
	var b = 10
	var c = a * b
	fmt.Println(c)

	//augmented assignment
	var i = 10
	i += 10 //i + 10 (10 + 10)
	fmt.Println(i)

	//unary operator
	i++ // i = i + 1
	fmt.Println(i)

	var negative = -100
	var positive = +100 //by default
	fmt.Println(negative)
	fmt.Println(positive)
}
