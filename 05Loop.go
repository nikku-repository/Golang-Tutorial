package main

import "fmt"

func main() {
	//Achive while loop using for keyword.....!
	num := 1
	for num <= 10 {
		fmt.Println(num * 5)
		num++
	}

	//for loop
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//Infinite for loop with break statement....!
	for {
		fmt.Println("nikku")
		break
	}

	//for loop with continue statement[if something match then continue, something else then print value].
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	//foreach loop
	for _, str := range "Nagesh" {
		fmt.Println(string(str))
	}
}
