package main

import "fmt"

func main() {
	var x string
	x = "first "
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	var a string = "hello"
	var y string = "world"
	fmt.Println(a == y)

	// short form

	short := "short"
	fmt.Println(short)
}