package main

import "fmt"

func random() any {
	return "Ups"
}

func main() {
	var result interface{} = random()
	//var resultString string = result.(string)
	//var resultInt int64 = result.(int64)
	//fmt.Println(resultString)
	//fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
