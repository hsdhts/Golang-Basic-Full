package main

import "fmt"

func main() {
	var nilaiUjian = 80
	var nilaiAbsensi = 80

	var lulusUjian bool = nilaiUjian >= 80
	var lulusAbsensi bool = nilaiAbsensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)
}
