package main

import "fmt"

type SomeStruct struct {
	someProperty int
}

func New(someProperty int) SomeStruct {
	return SomeStruct{someProperty: someProperty}
}

func main() {
	someStruct := New(1)
	fmt.Println(someStruct.someProperty)
}
