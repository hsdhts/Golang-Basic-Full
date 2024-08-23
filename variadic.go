package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, numbers := range numbers {
		total += numbers
	}
	return total

}

func main() {
	total := sumAll(10, 10, 10, 20, 30)
	fmt.Println(total)
}
