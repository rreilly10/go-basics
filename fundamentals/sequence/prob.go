
package main

import "fmt"
import "errors"

func main() {
	slice := make([]int, 3, 9)
	fmt.Println(slice)

	x := [6]string{"a","b","c","d","e","f"}
	fmt.Println(x[2:5])

	li := []int{
		48,96,86,68,
		57,82,63,70,
		37,34,83,27,
		19,97, 9,17,
	}

	min, err := smallest(li)

	if err == nil {
		fmt.Println(min)
	} else {
		fmt.Println(err)
	}

}


func smallest(set []int) (min int, e error) {
	if len(set) == 0 {
		return 0, errors.New("No elements detected in list")
	}
	
	min = set[0]
	for _, value := range set {
		if value < min {
			min = value
		}
	}

	return min, nil
}
