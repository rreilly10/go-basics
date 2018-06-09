package main

import "fmt"

func main () {
	x := [5]int{ 
		98, 
		93, 
		77, 
		82, 
		83,
	}

	total := 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}

	total = 0 

	for _, value := range x {
		total += value
	}

	avg := total/len(x)
	fmt.Println(avg)

}