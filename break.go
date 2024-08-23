package main

import "fmt"

func main() {
	for nilai := 0; nilai < 10; nilai++ {
		fmt.Println("Perulangan ke", nilai)

		if nilai == 5 {
			break
		}
	}

	// Todo: Mencetak angka ganjil
	for angka := 1; angka < 20; angka++ {
		if angka%2 == 0 {
			continue
		}

		fmt.Println(angka)
	}
}
