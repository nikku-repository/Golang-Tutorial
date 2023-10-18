package main

import (
	"fmt"
)

func main() {
	//initialize value into arr1 after declare arr1.
	var arr1 [5]int
	arr1[0], arr1[1], arr1[2], arr1[3], arr1[4] = 1, 2, 3, 4, 5
	fmt.Println(arr1)

	//declare and initialize an array in one line.
	var arr2 = [5]int{10, 20, 30, 40, 50}
	fmt.Println(arr2)

	//using shorthand declaration and initialization
	arr3 := [5]int{100, 101, 102, 103, 104}
	fmt.Println(arr3)

	//variadic array
	arr4 := [...]int{11, 12, 13, 14, 15}
	fmt.Println(arr4)

	//Two Dimensional array
	arr5 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr5)
}
