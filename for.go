package main

import "fmt"

func main() {
	for counter := 1; counter < 10; counter++ {
		fmt.Println("Perulangan ke:", counter)
	}

	names := []string{"Satu", "Dua", "Tiga"}
	for _, name := range names {
		fmt.Println(name)
	}
}
