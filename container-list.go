package main

import (
	"container/list"
	"fmt"
)

func main() {

	data := list.New()

	data.PushBack("Husada")
	data.PushBack("Karunia")
	data.PushBack("Hutasoit")
	data.PushFront("Budi")

	for element := data.Back(); element != nil; element.Prev() {
		fmt.Println(element.Value)
	}
}
