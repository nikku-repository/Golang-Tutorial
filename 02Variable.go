package main

import "fmt"

func main() {
	//variable declared using var keyword
	var a = "Nagesh"
	fmt.Println(a)

	//You can declare multiple variables at once
	var b, c int = 15, 12
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	//Variables declared without a corresponding initialization are zero-valued
	var e int
	fmt.Println(e)

	//:= syntax is shorthand for declaring and initializing a variable
	f := "code with nikku"
	fmt.Println(f)
}
