package main

import (
	"fmt"
	"golang-basic/database"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
