package main

import "fmt"

func main(){
	arr := [5]float64{1,2,3,4,5}
	x := arr[0:2]
	fmt.Println(x)

	// APPEND
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5,)
	fmt.Println(slice1, slice2)

	// COPY

	slice1 = []int{1,2,3}
	slice2 = make([]int, 2)
	
	fmt.Println("Copy")
	fmt.Println(slice1, slice2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}