package main

import "fmt"

func main() {
	var x float64 = 20.0
	fmt.Println(x)
	fmt.Printf("x is of type %T\n", x)

	// Late Initialization
	var num int
	var amount float32
	var isValid bool
	var name string
	num = 20
	amount = 49.99
	isValid = true
	name = "Bappy"
	fmt.Println(num, amount, isValid, name)
}
