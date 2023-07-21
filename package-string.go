package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Black Mamba", "Black"))
	fmt.Println(strings.Split("Husada Hutasoit", ""))
	fmt.Println(strings.ToLower("Husada Hutasoit"))
	fmt.Println(strings.ToUpper("husada hutasoit"))
	fmt.Println(strings.Trim("    Husada Hutasoit    ", " "))
	fmt.Println(strings.ReplaceAll("Husada Husada Husada Husada", "Husada", "Budi"))
}
