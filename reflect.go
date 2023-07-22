package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required: "true" max: "10"`
}

func main() {
	sample := Sample{"Husada"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
}
