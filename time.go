package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()

	fmt.Println(now)
	fmt.Println(now.YearDay())
	fmt.Println(now.Month())
	fmt.Println(now.Day())

	utc := time.Date(2020, 10, 10, 10, 10, 10, 10, time.UTC)

	fmt.Println(utc)

	layout := "2020-01-24"

	parse, _ := time.Parse(layout, "2020-01-23")
	fmt.Println(parse)

}
