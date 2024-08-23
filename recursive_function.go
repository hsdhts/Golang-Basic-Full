package main

import (
	"fmt"
)

func factorialLoop(value int64) int64 {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// Menggunakan Teknik Recursive
func recursiveFactrialLoop(value int64) int64 {
	if value == 1 {
		return 1
	} else {
		return value * recursiveFactrialLoop(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))

	fmt.Println(recursiveFactrialLoop(10))
}
