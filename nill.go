package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = NewMap("husada")

	if person == nil {
		fmt.Println("data masih kosong")
	} else {
		fmt.Println(person["name"])
	}
}
