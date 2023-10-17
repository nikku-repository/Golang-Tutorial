package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	const thita = 90
	fmt.Println(math.Sin(thita))

	//There are two type of constant in golang
	//1. Typed Constant
	//2. Untyped Constant
	const untypedInteger = 123
	const typedInteger = 123
	const untypedFloating = 123.12
	const typedFloatingPoint float64 = 123.12

	fmt.Println(untypedInteger)
	fmt.Println(typedInteger)
	fmt.Println(untypedFloating)
	fmt.Println(typedFloatingPoint)
}
