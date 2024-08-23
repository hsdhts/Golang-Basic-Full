package main

import "fmt"

func int(a int64, b int64) int64 {
	result := a + b
	return result
}

func main() {
	sum := int(3, 5)
	fmt.Println(sum)
}
