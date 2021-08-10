package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[3] = 20
	fmt.Println(arr, len(arr))

	arr1 := [3]int{1,2,3}
	fmt.Println(arr1)

	var arr2 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr2[i][j] = i + j
		}
	}
	fmt.Println(arr2)
}