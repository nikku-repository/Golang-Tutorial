package main

import (
	"bytes"
	"fmt"
	"reflect"
	"slices"
	"sort"
)

func main() {
	//declaring a slice <slice0> using var keyword and initialize with value in one line.
	var slice0 = []int{1, 2, 3, 4, 5}
	fmt.Println(slice0)

	//declaring a slice <slice1> using shorthand method and initialize with value in one line.
	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Println(slice1)

	//create a slice <slice2> with the help of slice1[2:4] [including index(2) and excluding index(4)]
	slice2 := slice1[2:4]
	fmt.Println(slice2)

	//create a slice <mySlice> using make function and print the lenght and cap of <mySlice>
	//and copy element of slice2 into mySlice
	var mySlice = make([]int, 5, 7)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	copy(mySlice, slice2)
	fmt.Println(mySlice)

	//create nil slice
	//there are two way to create a nil slice
	//1. var slice []int
	//2. slice := []int(nil)
	var slice4 []int
	if slice4 == nil {
		fmt.Println("This is nil slice..........!")
	}

	//compare two slice using slices.Equal method
	if slices.Equal(slice0, slice1) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//compare two slice using reflect.DeepEqual mathod

	if reflect.DeepEqual(slice0, slice1) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//append value into slice1
	slice1 = append(slice1, 60)

	//sorting slice using sort
	sort.Ints(slice1) // or sort.Strings for string slice

	//check IntsAreSorted
	if sort.IntsAreSorted(slice1) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	//trim a slice of bytes
	trimSlice := []byte{'!', '!', 'N', 'A', 'G', 'E', 'S', 'H', '#', '#'}
	fmt.Println(string(bytes.Trim(trimSlice, "!#")))

	//split a slice of bytes
	splitSlice := []byte{'!', '!', 'N', 'A', 'G', 'E', 'S', 'H', '#', '#'}

	fmt.Println(string(bytes.Split(splitSlice, []byte("AG"))[0]))
	fmt.Println(string(bytes.Split(splitSlice, []byte("AG"))[1]))

	//check for specific value contains into slice or not
	fmt.Println(slices.Contains(slice1, 20))

	//find max value from slice
	fmt.Println(slices.Max(slice1))

	//find min value from slice
	fmt.Println(slices.Min(slice1))
}
