package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error: ", err.Error())
	}

	valueInt, _ := strconv.Atoi("10000")
	fmt.Println(valueInt)
}
