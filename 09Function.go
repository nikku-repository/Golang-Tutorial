package main

import "fmt"

//Function parameters, return values together with types is called function signature.

//Function nature take nothing and return nothing
func nikku1() {
	fmt.Println("Welcome in nikku repo function....!")
}

//Function nature take nothing and return something
func nikku2() string {
	return "Welcome in nikku repo function....!"
}

//Function nature take something and return nothing
func nikku3(s string) {
	fmt.Println(s)
}

//Function nature take something and return something
func nikku4(a int, b int) int {
	return a + b
}

func main() {
	nikku1()
	fmt.Println(nikku2())
	nikku3("Welcome in nikku repo function....!")
	fmt.Println(nikku4(10, 20))
}
