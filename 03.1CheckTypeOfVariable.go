package main

import (
	"fmt"
	"reflect"
)

func main() {
	//check the datatype of variable
	a := "Nagesh"
	fmt.Printf("Data type of variable a\t:\t%T\n", a)
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a).Kind())

	//print the initial value of variable
	b := 10
	fmt.Printf("Initial value of b\t:\t%v\n", b)
	fmt.Println(reflect.ValueOf(b))
}
