package main

import (
	"fmt"
)

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println(filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	}
	return name
}

func main() {
	sayHelloWithFilter("Husada", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}
