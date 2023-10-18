package main

import "fmt"

func main() {

	var name = "Nagesh"
	if name == "Nagesh" {
		fmt.Println("This is Nagesh GitHub repository.......!")
	} else {
		fmt.Println("Print Else Message if condition Not true.......!")
	}

	var age = 16

	if age <= 16 {
		fmt.Println("You are under the age..........!")
	} else if age >= 18 {
		fmt.Println("You are eligible for vote........!")
	} else {
		fmt.Println("Print Else Message if condition Not Match.......!")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
