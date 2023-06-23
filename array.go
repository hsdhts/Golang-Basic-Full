package main

import "fmt"

func main() {

	var names [3]string
	names[0] = "Husada"
	names[1] = "Karunia"
	names[2] = "Hutasoit"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		95,
		90,
		80,
	}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi = [3]int{
		1,
		2,
		3,
	}
	fmt.Println(lagi)
	fmt.Println(lagi[0])
	fmt.Println(lagi[1])
	fmt.Println(lagi[2])

}
